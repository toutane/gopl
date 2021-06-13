package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/toutane/gopl/ch4/ex4.11/cli"
	"github.com/toutane/gopl/ch4/ex4.11/util"
)

// ClientID of gitool-app.
const ClientID = "Iv1.e2fe3805eaf6b9a2"

// Body of the request that requests device code and user verification
// codes from GitHub.
type Device struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
}

// Body of hte request that check if the user authorized the device.
type Access struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Status fuction retrun a message about the status of the authotification of
// the user.
func Status() string {
	hosts, err := util.LoadHosts()
	if err != nil {
		return fmt.Sprintf("Auth status: %s", err)
	}
	if hosts != nil {
		return fmt.Sprintf("\nYou are logged in to github.com as %s.", hosts.GitHubUser)
	}
	return fmt.Sprintf("\nYou are not logged in.")
}

// Login function log in to GitHub the user.
func Login() (string, error) {
	// Check if the user is already logged in.
	if IsLogged() {
		return fmt.Sprintf("You are already logged in."), nil
	}

	// Get the user GitHub username.
	err := GetUsername()
	if err != nil {
		return fmt.Sprint("Failed to get username."), err
	}

	// Get the device token and the user verifications codes.
	device, err := GetDevice()
	if err != nil {
		return fmt.Sprint("Failed to ask the device code."), err
	}

	fmt.Printf("\nHere are your authorization codes: %s\n", device.UserCode)

	// Copy to clipboard the user verification codes.
	cli.Clip(device.UserCode)

	// Open the user verification url in browser.
	exec.Command("open", device.VerificationURI).Start()

	// Ask the user if he enter his verification codes.
	isConfirm, err := cli.AskForConfirm("\nHave you authorized GiTool app")
	if err != nil {
		return fmt.Sprint("Failed to ask confirmation"), err
	}

	// Exit if user don't enter his verification codes.
	if !isConfirm {
		return fmt.Sprint("Please enter your codes."), fmt.Errorf("\nYou must authorize GiTool app.")
	}

	// Get the access token.
	access, err := GetAccess(device.DeviceCode)
	if err != nil {
		return fmt.Sprintf("Failed to ask the access token."), err
	}

	// Check if the token wa found.
	if access.AccessToken == "" {
		return fmt.Sprintf("No access token found."), fmt.Errorf("Any access token found.")
	}

	// Write access token to config file (app.env).
	util.Write("GITHUB_ACCESS_TOKEN", access.AccessToken)

	// Check if the authentification succeeded.
	if !IsLogged() {
		return fmt.Sprint("Failed write token."), fmt.Errorf("Failed to write token.")
	}

	// Return a message about the status of authentification.
	return Status(), nil
}

// GetUsername function prompt the user to enter his GitHub username.
func GetUsername() error {
	var username string
	fmt.Print("\nPlease enter your GitHub username: ")
	_, err := fmt.Scan(&username)
	util.Write("GITHUB_USER", username)
	return err
}

// GetAccess function returns the access token.
func GetAccess(DeviceCode string) (*Access, error) {
	params := map[string]string{"client_id": ClientID, "device_code": DeviceCode, "grant_type": "urn:ietf:params:oauth:grant-type:device_code"}
	url := "https://github.com/login/oauth/access_token?"

	for k, v := range params {
		url += k + "=" + v + "&"
	}
	url = url[:len(url)-1]

	client := http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result Access
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

// GetDevice function returns the device token and user verification codes.
func GetDevice() (*Device, error) {
	url := "https://github.com/login/device/code?client_id=" + ClientID
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result Device
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

// Logout function set to empty GitHubUser and GitHubAccessToken.
func Logout(path string) (string, error) {
	isLogged := IsLogged()
	util.Reset(path)
	if isLogged {
		return fmt.Sprintf("\nYou have been logged out."), nil
	}
	return fmt.Sprintf("\nYou are not logged in."), nil
}

// IsLogged function returns a boolean about the authentification status of user.
func IsLogged() bool {
	hosts, err := util.LoadHosts()
	if err != nil {
		return false
	}

	if hosts != nil {
		return true
	}
	return false
}

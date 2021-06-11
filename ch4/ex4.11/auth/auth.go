package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/toutane/gopl/ch4/ex4.11/cli"
	"github.com/toutane/gopl/ch4/ex4.11/util"
)

const ClientID = "Iv1.e2fe3805eaf6b9a2"

type Device struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
}

type Access struct {
	AccessToken string `json:"access_token"`
}

func Status() string {
	hosts, err := util.LoadHosts()
	if err != nil {
		return fmt.Sprintf("Auth status: %s", err)
	}
	if hosts != nil {
		return fmt.Sprintf("Logged in to github.com as %s.", hosts.GitHubUser)
	}
	return fmt.Sprintf("Not logged.")
}

func Login() (string, error) {
	if IsLogged() {
		return fmt.Sprintf("You are already logged in !"), nil
	}

	err := GetUsername()
	if err != nil {
		return fmt.Sprint("Failed to get username."), err
	}

	device, err := GetDevice()
	if err != nil {
		return fmt.Sprint("Failed to ask the device code."), err
	}

	fmt.Printf("Your code is: %s\n", device.UserCode)
	exec.Command("open", device.VerificationURI).Start()

	isConfirm, err := cli.AskForConfirm("Do you enter codes")
	if err != nil {
		return fmt.Sprint("Failes to ask confirmation"), err
	}

	if !isConfirm {
		return fmt.Sprint("Please enter your codes."), fmt.Errorf("Codes not confirmed.")
	}

	access, err := GetAccess(device.DeviceCode)
	if err != nil {
		return fmt.Sprintf("Failed to ask the access token."), err
	}

	if access.AccessToken == "" {
		return fmt.Sprintf("No access token found."), fmt.Errorf("Any access token found.")
	}

	util.Write("GITHUB_ACCESS_TOKEN", access.AccessToken)

	if !IsLogged() {
		return fmt.Sprint("Failed write token."), fmt.Errorf("Failed to write token.")
	}

	return Status(), nil
}

func GetUsername() error {
	var username string
	fmt.Print("Enter GitHub username: ")
	_, err := fmt.Scan(&username)
	util.Write("GITHUB_USER", username)
	return err
}

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

func Logout(path string) (string, error) {
	util.Reset(path)
	return fmt.Sprintf("You have been log out."), nil
}

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

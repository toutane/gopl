package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/toutane/gopl/ch4/ex4.11/util"
)

const ClientID = "Iv1.1ae6ec7b4bd67430"

type Device struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
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

	codes, err := GetDevice()
	if err != nil {
		return fmt.Sprintf("Failed to ask the device code..."), err
	}

	fmt.Printf("Your code is: %s\n", codes.UserCode)
	exec.Command("open", codes.VerificationURI).Start()
	return "", nil
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
	return fmt.Sprintf("You have been log out"), nil
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

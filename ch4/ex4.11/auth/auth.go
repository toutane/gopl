package auth

import (
	"fmt"

	"github.com/toutane/gopl/ch4/ex4.11/util"
)

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
	return "", nil
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

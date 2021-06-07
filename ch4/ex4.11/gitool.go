// Gitool let the user create, read, update and close GitHub issues.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/toutane/gopl/ch4/ex4.11/api"
)

func main() {
	/*
		data, err := cli.GetInputFromEditor()
		if err != nil {
			fmt.Printf("Fail at getting input from editor: %s\n", err)
		}
		fmt.Printf("%s\n", data)
	*/
	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	issueNumber := readCmd.Int("number", -1, "Number of the issue you want to read.")
	username := readCmd.String("username", "username", "GitHub username.")
	repo := readCmd.String("repo", "repo", "GitHub repository.")
	state := readCmd.String("state", "open", "Issue state.")

	if len(os.Args) < 2 {
		quit(nil)
	}

	switch os.Args[1] {

	case "read":
		readCmd.Parse(os.Args[2:])
		read(*username, *repo, *state, *issueNumber)

	default:
		quit(nil)
	}
}

// Read function prints one or all issues of a GitHub repo.
func read(owner, repo, state string, issueNumber int) {
	params := map[string]string{"state": state}
	if issueNumber < 0 {
		result, err := api.GetIssues(owner, repo, params)
		if err != nil {
			quit(err)
		}
		for _, item := range *result {
			fmt.Printf("%-5s #%-5d %9.9s %.55s\n", item.State, item.Number, item.User.Login, item.Title)
		}
	} else {
		result, err := api.GetIssue(owner, repo, params, issueNumber)
		if err != nil {
			quit(err)
		}
		item := *result
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func quit(err error) {
	if err != nil {
		fmt.Printf("Gitool: %s\n", err)
	}
	os.Exit(1)
}

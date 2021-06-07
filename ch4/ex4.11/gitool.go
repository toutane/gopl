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

	if len(os.Args) < 2 {
		quit(nil)
	}

	switch os.Args[1] {

	case "read":
		readCmd.Parse(os.Args[2:])
		read(*username, *repo, *issueNumber)

	default:
		quit(nil)
	}
}

// Read function prints one or all issues of a GitHub repo.
func read(owner, repo string, issueNumber int) {
	if issueNumber < 0 {
		result, err := api.GetIssues(owner, repo)
		if err != nil {
			quit(err)
		}
		for _, item := range *result {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	} else {
		result, err := api.GetIssue(owner, repo, issueNumber)
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

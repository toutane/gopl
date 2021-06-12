// Gitool let the user create, read, update and close GitHub issues.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/toutane/gopl/ch4/ex4.11/api"
	"github.com/toutane/gopl/ch4/ex4.11/auth"
	"github.com/toutane/gopl/ch4/ex4.11/cli"
	"github.com/toutane/gopl/ch4/ex4.11/util"
)

var username = "username"

func main() {

	hosts, err := util.LoadHosts()
	if err != nil {
		quit(err)
	}
	if hosts != nil {
		username = hosts.GitHubUser
	}

	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	readRepo := readCmd.String("repo", "repo", "Issue's repository.")
	readUsername := readCmd.String("username", username, "Issue's owner.")
	issueNumber := readCmd.Int("number", -1, "Number of the issue you want to read.")
	state := readCmd.String("state", "open", "Issues state.")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	title := createCmd.String("title", "New issue by ", "New issue's title.")
	createRepo := createCmd.String("repo", "repo", "New issue's repository.")

	closedCmd := flag.NewFlagSet("close", flag.ExitOnError)
	closedRepo := closedCmd.String("repo", "repo", "Issue's repository.")
	closedNumber := closedCmd.Int("number", -1, "Number of the issue you want to close.")

	updateCmd := flag.NewFlagSet("close", flag.ExitOnError)
	updateRepo := updateCmd.String("repo", "repo", "Issue's repository.")
	updateNumber := updateCmd.Int("number", -1, "Number of the issue you want to update.")

	authCmd := flag.NewFlagSet("auth", flag.ExitOnError)

	if !util.IsInitialized() {
		fmt.Println("CONFIG IS NOT INITIALIZED")
		err := util.InitializeConfig()
		if err != nil {
			quit(err)
		}
	}

	if len(os.Args) < 2 {
		fmt.Println(help("help"))
		quit(nil)
	}

	switch os.Args[1] {

	case "auth":
		if len(os.Args) < 3 {
			fmt.Println(help("auth"))
			quit(nil)
		}

		switch os.Args[2] {

		case "status":
			fmt.Println(auth.Status())

		case "login":
			mess, err := auth.Login()
			if err != nil {
				quit(err)
			}

			fmt.Println(mess)

		case "logout":
			mess, err := auth.Logout(".")
			if err != nil {
				quit(err)
			}

			fmt.Println(mess)

		default:
			fmt.Println(help("auth"))
			quit(nil)
		}
		authCmd.Parse(os.Args[3:])

	case "read":
		if len(os.Args) < 3 {
			fmt.Println(help("read"))
			quit(nil)
		}
		readCmd.Parse(os.Args[2:])
		read(*readUsername, *readRepo, *state, *issueNumber)

	case "create":
		if len(os.Args) < 3 {
			fmt.Println(help("create"))
			quit(nil)
		}
		createCmd.Parse(os.Args[2:])
		create(*createRepo, *title)

	case "update":
		if len(os.Args) < 3 {
			fmt.Println(help("update"))
			quit(nil)
		}
		updateCmd.Parse(os.Args[2:])
		if *updateNumber < 0 {
			fmt.Println(help("update"))
			quit(nil)
		}
		update(*updateRepo, *updateNumber)

	case "close":
		if len(os.Args) < 3 {
			fmt.Println(help("close"))
			quit(nil)
		}
		closedCmd.Parse(os.Args[2:])
		if *closedNumber < 0 {
			fmt.Println(help("close"))
			quit(nil)
		}
		closed(*closedRepo, *closedNumber)

	default:
		fmt.Println(help("help"))
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

// Create function creates a new issue.
func create(repo, title string) {
	if title == "New issue by " {
		title += username + "."
	}

	result, err := api.CreateIssue(repo, title)
	if err != nil {
		quit(err)
	}

	item := *result
	fmt.Printf("\n\nIssue number #%d successfully created by %s.\n", item.Number, item.User.Login)
}

// Update function update issue.
func update(repo string, number int) {
	result, err := api.UpdateIssue(repo, number)
	if err != nil {
		quit(err)
	}

	item := *result
	fmt.Printf("\n\nIssue number #%d successfully updated by %s.\n", item.Number, item.User.Login)
}

// Closed function close issue.
func closed(repo string, number int) {
	result, err := api.ClosedIssue(repo, number)
	if err != nil {
		quit(err)
	}

	item := *result
	fmt.Printf("\n\nIssue number #%d successfully closed by %s.\n", item.Number, item.User.Login)
}

// Help function prints help message for each command.
func help(command string) string {
	return cli.HelpMessages[command]
}

// Quit function prints a error and exit.
func quit(err error) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}

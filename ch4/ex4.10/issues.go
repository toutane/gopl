// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/toutane/gopl/ch4/ex4.10/github"
)

func main() {
	now := time.Now()
	var lessThanMonth, lessThanYear, older []github.Issue
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n\n", result.TotalCount)
	for _, item := range result.Items {
		switch {
		case !item.CreatedAt.Before(now.AddDate(0, -1, 0)):
			lessThanMonth = append(lessThanMonth, *item)
		case !item.CreatedAt.Before(now.AddDate(-1, 0, 0)):
			lessThanYear = append(lessThanYear, *item)
		default:
			older = append(older, *item)
		}
	}
	if len(lessThanMonth) > 0 {
		fmt.Println("Less than a month")
		for _, item := range lessThanMonth {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	if len(lessThanYear) > 0 {
		fmt.Println("\nLess than a year")
		for _, item := range lessThanYear {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}

	}
	if len(older) > 0 {
		fmt.Println("\nOlder than a year")
		for _, item := range older {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}

	}
}

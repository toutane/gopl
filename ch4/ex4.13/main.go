package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/toutane/gopl/ch4/ex4.13/movies"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchPage := searchCmd.Int("p", 1, "The page number of the results.")
	searchKeywords := searchCmd.String("k", "indiana jones", "The keywords of the movie you are looking for (separeted by a space)")

	posterCmd := flag.NewFlagSet("poster", flag.ExitOnError)
	posterMovieId := posterCmd.Int("id", 1, "The id of the movie you want the poster of.")

	if len(os.Args) < 2 {
		return errors.New("missing command")
	}

	command := os.Args[1]
	switch command {

	case "search":
		searchCmd.Parse(os.Args[2:])

		keywords := strings.Split(*searchKeywords, " ")
		page := *searchPage
		if page < 0 {
			page = 1
		}
		if err := movies.Search(keywords, page); err != nil {
			return fmt.Errorf("search failed: %s", err)
		}

	case "poster":
		posterCmd.Parse(os.Args[2:])

		movieId := *posterMovieId
		if movieId < 0 {
			movieId = 1
		}
		if err := movies.Poster(movieId); err != nil {
			return fmt.Errorf("poster failed: %s", err)
		}

	default:
		return errors.New("wrong command (search or poster)")
	}
	return nil
}

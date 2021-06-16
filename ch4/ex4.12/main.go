package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/toutane/gopl/ch4/ex4.12/xkcd"
)

func main() {
	populateCommand := flag.NewFlagSet("populate", flag.ExitOnError)
	lastNumber := populateCommand.Int("n", 1, "Number of comics you want to download for your index")

	if len(os.Args) < 2 {
		usageDie()
	}

	switch os.Args[1] {
	case "populate":
		populateCommand.Parse(os.Args[2:])
		if *lastNumber < 1 || *lastNumber > 2476 {
			usageDie()
		}

		err := xkcd.Populate(*lastNumber)
		if err != nil {
			exit(err)
		}
		fmt.Println("Finish to populate.")
	case "search":
		if len(os.Args) < 3 {
			usageDie()
		}
		keywords := os.Args[2:]

		foundComics, err := xkcd.Search(keywords)
		if err != nil {
			exit(err)
		}

		comicsN := len(*foundComics)
		if comicsN == 0 {
			fmt.Println("\nNone of the comics in the indew match the keywords provided.")
			fmt.Println("\nTry ./main populate -n=50 to download the first 50 comics of xkcd.com in to indew.")
			return
		}
		fmt.Printf("\n%d comics in the index match the keywords provided:", comicsN)
		fmt.Println()
		for _, comic := range *foundComics {
			fmt.Println("\n-------------------------------------------------------\n")
			fmt.Println(comic.Title)
			fmt.Println()
			fmt.Println(comic.Transcript)
		}
	default:
		usageDie()
	}

}

func usageDie() {
	const message = `
  USAGE
    
    ./xkcd command [flags]

  CORE COMMANDS
    
    populate     Download comics to create an index.
      -n           Number of the last comic to download (required; must be between 1 and 2476).
    search       Search commics in to the index.
    `
	fmt.Println(message)
	os.Exit(1)
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

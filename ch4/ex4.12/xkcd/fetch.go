package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Comic struct {
	Title      string
	Transcript string
	Image      string `json:"img"`
	Number     int    `json:"num"`
}

func readIndex() (*[]Comic, error) {
	var result []Comic

	file, err := os.OpenFile("index.json", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return &result, nil
	}

	if err := json.NewDecoder(file).Decode(&result); err != nil {
		file.Close()
		return &result, err
	}
	file.Close()
	return &result, nil
}

func getLastComicNumber() (int, error) {
	comics, err := readIndex()
	if err != nil {
		return 1, err
	}

	arr := *comics
	if len(arr) == 0 {
		return 1, nil
	}

	number := arr[len(arr)-1].Number
	return number + 1, nil
}

func get(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var result Comic

	//Decode resp.Body.
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func writeIndex(comics *[]Comic) error {
	oldComics, err := readIndex()
	if err != nil {
		fmt.Println("An error occures when try to read the index")
		return err
	}

	file, err := os.OpenFile("index.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	var newComics []Comic
	for _, comic := range *oldComics {
		newComics = append(newComics, comic)
	}
	for _, comic := range *comics {
		newComics = append(newComics, comic)
	}

	err = os.Truncate("index.json", 0)
	if err != nil {
		fmt.Println("An error occures when try to truncate the file")
		return err
	}

	if err := json.NewEncoder(file).Encode(newComics); err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil

}

func Populate(lastNum int) error {
	fmt.Println("Starting to populate database...")

	var comics []Comic

	//Get the number of the last comic.
	startNumber, err := getLastComicNumber()
	if err != nil {
		return err
	}
	fmt.Printf("Start populating index from comic number %d.\n", startNumber)

	//Fetching from comic number 1 to lastNum.
	for i := startNumber; i <= lastNum; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		//Get the comic number i.
		comic, err := get(url)
		if err != nil {
			return err
		}
		comics = append(comics, *comic)
	}

	//Write the array of comics in to index.json.
	err = writeIndex(&comics)
	if err != nil {
		fmt.Println("Error when write in index")
		return err
	}
	return nil
}

func Search(keywords []string) (*[]Comic, error) {
	var matches []Comic
	//Get all the comics there are in the index.
	addrIndexComics, err := readIndex()
	if err != nil {
		return &matches, err
	}
	indexComics := *addrIndexComics
	if len(indexComics) == 0 {
		return &matches, fmt.Errorf("xkcd : index is empty, could'nt search in it.")
	}

	//Iterate on each comic and see if contains keywords.
	for _, comic := range indexComics {
		isInMatches := false
		for _, word := range keywords {
			if strings.Contains(comic.Title, word) || strings.Contains(comic.Transcript, word) && !isInMatches {
				matches = append(matches, comic)
				isInMatches = true
			}
		}
	}

	return &matches, nil

}

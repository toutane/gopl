package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const APIURL = "https://api.themoviedb.org/3"

type MovieList struct {
	Page         int
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Movies       []Movie `json:"results"`
}

type Movie struct {
	Id    int
	Title string
}

func Search(keywords []string, page int) error {
	movieList, err := searchMoviesByKeywords(keywords, page)
	if err != nil {
		return fmt.Errorf("poster command failed: %s", err)
	}

	if movieList.Page > movieList.TotalPages {
		return fmt.Errorf("You are far too greedy... (Try search -p=%d -k='%s')", movieList.TotalPages, strings.Join(keywords, " "))
	}

	fmt.Printf("\n%d movies match with keywords. Page %d/%d\n\n", movieList.TotalResults, movieList.Page, movieList.TotalPages)
	fmt.Println("ID\t  Title")
	for _, movie := range movieList.Movies {
		fmt.Printf("%-9d %.55s\n", movie.Id, movie.Title)
	}
	return nil
}

func searchMoviesByKeywords(keywords []string, page int) (*MovieList, error) {
	url := APIURL + "/search/movie?api_key=" + APIKEY + "&query=" + strings.Join(keywords, "%20") + "&page=" + strconv.Itoa(page)

	//Do the fetch.
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("search fetch failed: %s", err)
	}

	//Decode response body.
	var result MovieList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("json decode failed: %s", err)
	}

	resp.Body.Close()
	return &result, nil
}

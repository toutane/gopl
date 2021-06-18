package movies

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

type ImageList struct {
	Id     int
	Images []Image `json:"posters"`
}

type Image struct {
	FilePath string `json:"file_path"`
}

func Poster(movieId int) error {
	imageList, err := getImages(movieId)
	if err != nil {
		return fmt.Errorf("poster failed: %s", err)
	}

	//We keep the first only.
	//TODO: Allow user to chose the poster he wants.
	poster := imageList.Images[0]

	if err := downloadPoster(movieId, &poster); err != nil {
		return fmt.Errorf("poster failed: %s", err)
	}
	fmt.Println("The poster has been successfully downloaded.")
	return nil
}

func getImages(movieId int) (*ImageList, error) {
	url := APIURL + "/movie/" + strconv.Itoa(movieId) + "/images?api_key=" + APIKEY

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("images fetch failed: %s", err)
	}

	var result ImageList

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("json decode failed: %s", err)
	}

	resp.Body.Close()
	return &result, nil
}

func downloadPoster(movieId int, poster *Image) error {
	url := "https://themoviedb.org/t/p/w440_and_h660_face" + poster.FilePath
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("poster download failed: %s", err)
	}

	if err := writePoster(movieId, resp); err != nil {
		return fmt.Errorf("poster download failed: %s", err)
	}

	return nil
}

func writePoster(movieId int, response *http.Response) error {
	file, err := os.Create(strconv.Itoa(movieId) + ".jpg")
	if err != nil {
		return fmt.Errorf("create file failed: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("write poster failed: %s", err)
	}

	return nil
}

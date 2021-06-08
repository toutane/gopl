// Exercise1-10 writes in a file the times and sizes of specified URLs fetchs
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exercise1-10: %v\n", err)
		os.Exit(1)
	}
	f.WriteString(fmt.Sprintf("%v\n\n", time.Now()))
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		f.WriteString(<-ch)
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n\n", time.Since(start).Seconds()))
	f.WriteString("--------------------------------------------------------\n\n")
	f.Close()

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak ressources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

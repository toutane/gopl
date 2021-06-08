// Exercise1-9 prints the content found at each specified url and the status of
// the http request
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1-9: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "HTTP Status: %v", resp.Status)
		time.Sleep(3 * time.Second)
		n, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("%v bytes", n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1-9: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

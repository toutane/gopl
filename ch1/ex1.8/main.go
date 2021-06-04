// Exercise1-8 print the content for each specified URL and added http:// if it
// is missing
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1-8: %v\n", err)
			os.Exit(1)
		}
		n, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("%v bytes", n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1-8: url: %s %v\n", url, err)
			os.Exit(1)
		}
	}
}

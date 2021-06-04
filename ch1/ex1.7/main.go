// Exercise1-7 prints the content found at each specified URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1-7: %v\n", err)
			os.Exit(1)
		}
		n, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("%v bytes\n", n)
		resp.Body.Close() // I don't no if it is necessary
		if err != nil {
			fmt.Fprintf(os.Stderr, "exericse1-7: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

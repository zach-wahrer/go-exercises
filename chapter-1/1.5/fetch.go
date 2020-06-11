// Fetch prints the raw content found at a URL
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
			url = addPrefix(url)
		}
		response, getError := http.Get(url)
		if getError != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", getError)
			os.Exit(1)
		}
		fmt.Println("Status Code: ", response.Status)
		_, parseError := io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if parseError != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, parseError)
			os.Exit(1)
		}
	}
}

func addPrefix(url string) string {
	return "http://" + url
}

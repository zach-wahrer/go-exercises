// Fetch prints the raw content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, getError := http.Get(url)
		if getError != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", getError)
			os.Exit(1)
		}
		_, parseError := io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if parseError != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, parseError)
			os.Exit(1)
		}
	}
}

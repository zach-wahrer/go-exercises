// Fetchall fetches URLs in parallel and reports the times and sizes.
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
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	file := createFile()
	for range os.Args[1:] {
		_, err := file.WriteString(<-ch)
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func createFile() *os.File {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		file.Close()
	}
	return file
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

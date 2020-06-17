package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Download(title, URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s.jpg", title))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	file.Write(body)

	return nil
}

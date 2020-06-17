package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const OMDBURL = "http://www.omdbapi.com/?apikey="

func buildSearchURL(title string) string {
	APIKey := os.Getenv("OMDB_API_KEY")
	if APIKey == "" {
		log.Fatal("You must set an API key for OMDB - ie.: `export OMDB_API_KEY=your_key_here`")
	}
	return fmt.Sprintf("%s%s&t=%s", OMDBURL, APIKey, title)

}

func Search(title string) (*Movie, error) {
	q := url.QueryEscape(title)
	resp, err := http.Get(buildSearchURL(q))
	if err != nil {
		return nil, fmt.Errorf("search failed: %s", resp.Status)
	}
	var movie Movie

	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &movie, nil
}

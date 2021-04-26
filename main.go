package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status int
}

func main() {
	results := make(map[string]string)
	channel := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i:=0; i<len(urls); i++ {
		result := <- channel
		results[result.url] = fmt.Sprint(result.status)
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// send only
func hitURL(url string, channel chan<- requestResult)  {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("ERROR❌", url, err, resp.StatusCode)
	} 
	channel <- requestResult{url: url, status: resp.StatusCode}
}
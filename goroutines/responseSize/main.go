package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	pageURL  string
	pageSize int
}

func responseSize(url string, channel chan page) { // recieve url string and channel page struct
	fmt.Println("Getting URL:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- page{pageURL: url, pageSize: len(body)}
}

func main() {
	pages := make(chan page) //channel holds page struct
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}
	for _, url := range urls {
		go responseSize(url, pages) // pass the channel to responseSize
	}
	for i := 0; i < len(urls); i++ {
		pageDetails := <-pages
		fmt.Printf("%s: %d\n", pageDetails.pageURL, pageDetails.pageSize)
	}
}

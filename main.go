package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // blocking call because it takes some time to run
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is responding to traffic & is ok!")
}

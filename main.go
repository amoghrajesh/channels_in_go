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

	c := make(chan string)

	// this wont work because main function has higher weightage which goes off without caring if child routine
	// completed or not.
	for _, link := range links {
		go checkLink(link, c)
	}

	// receiving messages from a channel is a blocking call
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // blocking call because it takes some time to run
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is responding to traffic & is ok!")
	c <- link
}

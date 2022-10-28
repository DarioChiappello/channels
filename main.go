package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//Channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {

		// literal function
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		// Send message into a channel
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}

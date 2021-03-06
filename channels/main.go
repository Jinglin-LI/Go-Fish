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

	// create a channel
	c := make(chan string)

	for _, link := range links {
		// Create a new thread go routine and run this function with it
		go checkLink(link, c)
	}

	// infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// pause in the Main Routine
		// time.Sleep(5 * time.Second)
		// go checkLink(l, c)

		go func(link string) { // l vs link, Main Routine vs Child Routine
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // argument for the function literal, goes to the memory
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

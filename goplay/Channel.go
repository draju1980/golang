package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://github.com",
		"http://gitlab.com",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://techwithaly.ongraphy.com",
		"http://arriqaaq.substack.com",
	}
	// creatation of channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // enableing go routine and passing channel information
	}
	// for {
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		// loopint through channels //
		// go checkLink(l, c)
		go func(link string) { // function literal or lambda function
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // sending message to channel
		return
	}
	fmt.Println(link, "is up!")
	c <- link // sending message to channel
}

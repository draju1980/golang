package main

import (
	"fmt"
	"net/http"
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
	for _, link := range links {
		checkLink(link)
	}
}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

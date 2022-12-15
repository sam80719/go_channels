package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://go.dev/",
		"https://tw.yahoo.com/",
		"https://www.facebook.com/",
		"https://www.google.com/webhp?hl=zh-TW",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) { // js, php: anonmous function; go: function literal; python:lambda
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "connection failure")
		c <- link
		return
	}

	fmt.Println(link, "connection success")
	c <- link
}

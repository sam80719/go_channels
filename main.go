package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://go.dev/",
		"https://tw.yahoo.com/",
		"https://www.facebook.com/",
		"https://www.google.com/webhp?hl=zh-TW",
	}

	c := make(chan string) // 建立字串類型的channels

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(i, <-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "connection failure")
		c <- "connection failure!"
		return
	}

	fmt.Println(link, "connection success")
	c <- "connection success!"
}

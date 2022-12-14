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
	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "connection failure")
		return
	}

	fmt.Println(link, "connection success")
}

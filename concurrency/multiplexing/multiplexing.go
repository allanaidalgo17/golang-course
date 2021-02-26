package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getTitle(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func forward(source <-chan string, destiny chan string) {
	for {
		destiny <- <-source
	}
}

func multiplexing(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(input1, c)
	go forward(input2, c)
	return c
}

func main() {
	c := multiplexing(
		getTitle("https://www.google.com/", "https://github.com/"),
		getTitle("https://www.amazon.com.br/", "https://www.youtube.com/"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

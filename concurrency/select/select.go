package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func fasterOne(url1, url2, url3 string) string {
	c1 := getTitle(url1)
	c2 := getTitle(url2)
	c3 := getTitle(url3)

	// control structure for concurrency
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Second):
		return "None"
		// default:
		// 	return "No answer yet"
	}
}

func main() {
	faster := fasterOne(
		"https://www.amazon.com.br/",
		"https://www.youtube.com/",
		"https://www.google.com/",
	)

	fmt.Println(faster)
}

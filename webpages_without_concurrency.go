package main

import (
	"fmt"
	"net/http"
	"time"
)

var webpages = []string{
	"https://google.com",
	"https://github.com",
	"https://fb.com",
	"https://instagram.com",
	"https://go.dev",
	"https://stackoverflow.com",
	"https://reddit.com",
	"https://linkedin.com",
	"https://youtube.com",
	"https://twitter.com",
	"https://amazon.com",
	"https://medium.com",
	"https://wikipedia.org",
	"https://netflix.com",
	"https://quora.com",
	"https://apple.com",
	"https://microsoft.com",
	"https://openai.com",
	"https://coursera.org",
	"https://udemy.com",
}

func main() {
	start := time.Now()
	for i:=0;i<len(webpages);i++{
		getStatusCode(webpages[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %d ms\n", elapsed.Milliseconds())
}

func getStatusCode(webpage string) {
	res, err := http.Get(webpage)
	if err!=nil{
		fmt.Printf("Some problem in opening %s\n", webpage)
	}else{
		fmt.Printf("%d status code for %s\n", res.StatusCode, webpage)
	}
}
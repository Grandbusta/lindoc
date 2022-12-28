package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("lindoc running")
	res, err := http.Get("https://betterprogramming.pub/deep-dive-into-concurrency-of-go-93002344d37b")
	// res, err := http.Get("https://www.keypup.io/blog/3-reasons-why-software-developers-use-social-media")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	con := html.UnescapeString(string(content))
	fmt.Println(con)
	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.Write(content)
}

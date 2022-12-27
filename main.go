package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("lindoc running")
	res, err := http.Get("https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	// con := html.UnescapeString(string(content))
	fmt.Println(string(content))
}

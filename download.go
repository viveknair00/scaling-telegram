package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	c := make(chan bool)

	//this is where the magic happens
	for i := 1; i < 25; i++ {
		go GetUrl("https://jsonplaceholder.typicode.com/posts/", i, c)
		<-c
	}
}

func Add(x, y int) (res int) {
	return x + y
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUrl(url string, part int, c chan bool) {
	page := url + strconv.Itoa(part)
	resp, err := http.Get(page)

	log.Println(page)
	check(err)

	reader, err := io.ReadAll(resp.Body)

	check(err)

	log.Print(part)

	filePath := "/tmp/test_" + strconv.Itoa(part) + ".txt"
	bodyString := string(reader)
	f, err := os.Create(filePath)
	check(err)
	f.WriteString(bodyString)

	log.Print("Completed writing to", filePath)
	c <- true
}

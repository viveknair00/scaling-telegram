package main

import (
	"fmt"

	"github.com/viveknair00/scaling-telegram/download"
)

func main() {
	c := make(chan bool)

	//this is where the magic happens
	for i := 1; i < 25; i++ {
		go download.GetUrl("https://jsonplaceholder.typicode.com/posts/", i, c)
		<-c
	}

	fmt.Println(download.Add(6, 4))
	fmt.Println(download.Subtract(5, 4))
}

package main

import (
	"strconv"

	"github.com/viveknair00/scaling-telegram/download"
	"github.com/viveknair00/scaling-telegram/vk_consumer"
)

func main() {

	// exampleDownload()
	vk_consumer.GetSnapshotPageData()
}

func exampleDownload() {
	c := make(chan bool)

	//this is where the magic happens
	for i := 1; i < 25; i++ {
		page := "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i)
		filePath := "/tmp/test_" + strconv.Itoa(i) + ".txt"
		go download.GetUrl(page, filePath, c)
		<-c
	}
}

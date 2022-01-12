package download

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/viveknair00/scaling-telegram/utils"
)

func GetUrl(url string, filePath string, c chan bool) {
	FileWriter(filePath, ReadUrl(url))
	c <- true
}

func ReadUrl(url string) (reader []byte) {
	resp, err := http.Get(url)
	log.Println(url)
	utils.Check(err)
	reader1, err := io.ReadAll(resp.Body)
	utils.Check(err)
	return reader1
}

func FileWriter(filePath string, reader []byte) {
	bodyString := string(reader)
	f, err := os.Create(filePath)
	utils.Check(err)
	f.WriteString(bodyString)

	log.Print("Completed writing to", filePath)
}

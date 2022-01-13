package vk_consumer

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/viveknair00/scaling-telegram/download"
	"github.com/viveknair00/scaling-telegram/utils"
)

type page struct {
	brand  string
	market string
	name   string
	pageId string
}

type parameters struct {
	baseApiUrl            string
	apiVersion            string
	pageId                string
	urlDate               string
	token                 string
	statsMethodName       string
	postsMethodName       string
	videoMethodName       string
	subscribersMethodName string
}

var pages [5]page = [5]page{page{
	brand:  "Vogue",
	market: "Russia",
	name:   "Vogue Russia",
	pageId: "24396213",
},
	page{
		brand:  "GQ",
		market: "Russia",
		name:   "GQ Russia",
		pageId: "2089898",
	},
	page{
		brand:  "Tatler",
		market: "Russia",
		name:   "Tatler",
		pageId: "136270576",
	},
	page{
		brand:  "Glamour",
		market: "Russia",
		name:   "Glamour Russia",
		pageId: "11064089",
	},
	page{
		brand:  "AD",
		market: "Russia",
		name:   "Architecture Digest",
		pageId: "36420241",
	},
}

func buildUrl(params parameters, dataType string, interval int, offset int) (url string) {

	if dataType == "post" {

		s := fmt.Sprintln(params.baseApiUrl, params.postsMethodName, "owner_id=-", params.pageId, "&count=", interval, "&offset=", offset, "&access_token=", params.token, "&v=", params.apiVersion)
		url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
		return url

	} else if dataType == "video" {
		s := fmt.Sprintln(params.baseApiUrl, params.videoMethodName, "owner_id=-", params.pageId, "&count=", interval, "&offset=", offset, params.token, "&v=", params.apiVersion)
		url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
		return url

	} else {
		return ""
	}

}

func RetrieveData(params parameters, dataType string, c chan bool) {
	response := download.ReadUrl(buildUrl(params, dataType, 1, 0))
	count, err := jsonparser.GetInt(response, "response", "count")
	utils.Check(err)
	fmt.Println(count)

	pageSize := int64(100)

	for pageStart := int64(0); int64(pageStart) < count; pageStart += pageSize {
		fmt.Printf("value of a: %d\n", pageStart)
		filePath := "/tmp/posts/vk_" + "posts_" + params.pageId + "_" + strconv.Itoa(int(pageStart)) + ".json"
		go download.GetUrl(buildUrl(params, dataType, int(pageSize), int(pageStart)), filePath, c)
		time.Sleep(1 * time.Second)
	}
}

func GetSnapshotPageData() {

	videoAccessToken := "PUT_YOUR_TOKEN_HERE"

	fmt.Println("start")

	for idx, page := range pages {
		fmt.Println(idx, page)
		urlDate := fmt.Sprintf("&timestamp_from=%s&timestamp_to=%s", "1641859200", "1641945600")

		param := parameters{
			baseApiUrl:            "https://api.vk.com/method/",
			apiVersion:            "5.92",
			pageId:                page.pageId,
			urlDate:               urlDate,
			token:                 videoAccessToken,
			statsMethodName:       "stats.get?",
			postsMethodName:       "wall.get?",
			videoMethodName:       "video.get?",
			subscribersMethodName: "groups.getMembers?",
		}

		// go getStats(param, c)
		// go getPosts(param, c)
		// go getVideos(param, c)
		// go getSubscribers(param, c)
		c := make(chan bool)
		go RetrieveData(param, "post", c)
		<-c

	}

	fmt.Println("Finish")

}

func getPageStats(params parameters, c chan bool) {
	s := fmt.Sprintln(params.baseApiUrl, params.statsMethodName, "group_id=", params.pageId, params.urlDate, "&access_token=", params.token, "&v=", params.apiVersion)
	url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
	filePath := "/tmp/vk_" + "page_" + params.pageId
	download.GetUrl(url, filePath, c)
}

func getStats(params parameters, c chan bool) {
	s := fmt.Sprintln(params.baseApiUrl, params.statsMethodName, "group_id=", params.pageId, params.urlDate, "&access_token=", params.token, "&v=", params.apiVersion)
	url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
	filePath := "/tmp/vk_" + "stats_" + params.pageId
	download.GetUrl(url, filePath, c)
}

func getPosts(params parameters, c chan bool) {
	s := fmt.Sprintln(params.baseApiUrl, params.postsMethodName, "owner_id=-", params.pageId, "&count=1&offset=0", "&access_token=", params.token, "&v=", params.apiVersion)
	url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
	filePath := "/tmp/vk_" + "posts_" + params.pageId
	download.GetUrl(url, filePath, c)
}

func getVideos(params parameters, c chan bool) {
	s := fmt.Sprintln(params.baseApiUrl, params.videoMethodName, "owner_id=-", params.pageId, "&count=1&offset=0", "&access_token=", params.token, "&v=", params.apiVersion)
	url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
	filePath := "/tmp/vk_" + "videos_" + params.pageId
	download.GetUrl(url, filePath, c)
}

func getSubscribers(params parameters, c chan bool) {
	s := fmt.Sprintln(params.baseApiUrl, params.subscribersMethodName, "&count=1&offset=0", "&group_id=", params.pageId, "&access_token=", params.token, "&v=", params.apiVersion)
	url := strings.TrimSuffix(strings.ReplaceAll(s, " ", ""), "\n")
	filePath := "/tmp/vk_" + "subscribers_" + params.pageId
	download.GetUrl(url, filePath, c)
}

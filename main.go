package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// ItemInfo はレコードの項目
type ItemInfo struct {
	SiteID             int
	UpdatedTime        time.Time
	IssueID            int
	Ver                int
	Title              string
	Body               string
	StartTime          time.Time
	CompletionTime     time.Time
	WorkValue          float64
	ProgressRate       float64
	RemainingWorkValue float64
	Status             int
	Manager            int
	Owner              int
	Comments           []string
	Creator            int
	Updator            int
	CreatedTime        time.Time
	ItemTitle          string
	ClassHash          string
	NumHash            string
	DateHash           string
	DescriptionHash    string
	CheckHash          string
	AttachmentsHash    string
}

// Response は取得したデータに関する情報
type Response struct {
	Offset     int
	PageSize   int
	TotalCount int
	Data       *[]ItemInfo
}

// ResponseInfo はAPIで取得したレスポンス
type ResponseInfo struct {
	StatusCode     int
	LimitPerDate   int
	LimitRemaining int
	Response       *[]Response
}

func main() {
	os.Exit(run(os.Args))
}

type RequestBody struct {
	ApiVersion float64
	ApiKey     string
}

func run(args []string) int {
	client := &http.Client{}
	client.Timeout = time.Second * 15

	header := http.Header{}
	header.Add("Content-Type", "application/json")

	requestBody := RequestBody{
		ApiVersion: 1.1,
		ApiKey:     "XXXXX...",
	}

	json, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "https://pleasanter.net/fs/api/items/XXXXX/get", bytes.NewBuffer(json))
	if err != nil {
		return 1
	}
	req.Header = header

	res, err := client.Do(req)
	if err != nil {
		return 2
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 3
	}

	fmt.Printf("%s", body)

	return 0
}

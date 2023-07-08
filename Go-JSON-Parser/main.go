package main

import (
	"encoding/json"
	"fmt"
	apicall "go-json-parser/api-call"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	jsonData := `{
		"bill_id":"fba23-4444-sd-weeqw-q",
		"date":"21/12/2023",
		"price":1000,
		"items":[
			{"id":"adsases-serrr-sadsf"}
		],
		"channel":{
			"id":"adabs-222ae-dwe2-sd3eq"
		}
	}
	`
	url := "https://jsonplaceholder.typicode.com/posts"

	//ApiService global method of api-call package imported as apicall
	//ApiService declared in api-call/api-call.go
	//args : url string <URL of the server to be requested>
	//       &wg *sync.WaitGroup <Address of the wait group to add threads>
	go apicall.ApiService(url, &wg) //invoked as a go routine for mutithreaded implementation

	// Below method will run in main thread
	parseData, err := jsonParser(jsonData) //parse jsonData to the package level method jsonParser
	if err != nil {
		fmt.Printf("error occured: %v\n", err.Error())
	}
	fmt.Println(parseData)
	wg.Wait() //wait for all the threads in wait group to complete
}

func jsonParser(jsonData string) (map[string]interface{}, error) {
	var parseData map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &parseData)
	return parseData, err
}

package apicall

import (
	"io"
	"log"
	"net/http"
	"sync"
)

// pass wg by reference
func ApiService(url string, wg *sync.WaitGroup, channel chan string) (string, error) {
	var ret string
	var eret error
	wg.Add(1) //add thread to wait group
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("error occured:", err.Error())
		eret = err
	} else {
		body, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			log.Fatalln("error occured:", err.Error())
			eret = err1
		} else {
			ret = string(body)
			channel <- ret //push response into the channel
			// fmt.Printf("response fetched\nbody: %v", ret)
		}
	}
	wg.Done() //remove from wait group

	return ret, eret
}

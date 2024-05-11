package main

import (
	"NEOs/configs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Response struct {
	Total              interface{}        `json:"element_count"`
	Near_earth_objects Near_earth_objects `json:"near_earth_objects"`
}

type Near_earth_objects []struct {
	Date  interface{} `json:"date"`
	ID    interface{} `json:"id"`
	Name  interface{} `json:"name"`
	Check interface{} `json:"is_potentially_hazardous_asteroid"`
}

func NewResponse(total interface{}, near_earth_objects Near_earth_objects) *Response {
	return &Response{
		Total:              total,
		Near_earth_objects: near_earth_objects,
	}
}

var data map[string]interface{}

func main() {

	// Read Configs
	configs := configs.ReadConfigs()

	// Send GET request
	resp, err := http.Get(configs.Link)

	if err != nil {
		fmt.Println("No response from request")
	}

	// Read Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read a JSON due to %s", err)
	}

	// Decode a Json
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	// Assigning a Total value
	result := NewResponse(data["element_count"], Near_earth_objects{})

	// Creating a Wait Group and Mutex
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	// Iteration through data
	for date, value := range data["near_earth_objects"].(map[string]interface{}) {
		var id, name, check interface{}
		wg.Add(1)

		// Concurrent data fetching for each day
		go getDay(&id, &name, &check, value, &date, &wg, &mutex, result)
	}
	wg.Wait()

	// Returning a result
	jsonResult, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(string(jsonResult))
}

// Assigning received data to Response object
func addResponse(date *string, id, name, check *interface{}, wg *sync.WaitGroup, result *Response) {
	result.Near_earth_objects = append(result.Near_earth_objects, struct {
		Date  interface{} "json:\"date\""
		ID    interface{} "json:\"id\""
		Name  interface{} "json:\"name\""
		Check interface{} "json:\"is_potentially_hazardous_asteroid\""
	}{date, id, name, check})
}

// Fetching data from one day
func getDay(id, name, check *interface{}, value interface{}, date *string, wg *sync.WaitGroup, mutex *sync.Mutex, result *Response) {
	mutex.Lock()
	for _, element := range value.([]interface{}) {
		*id = element.(map[string]interface{})["id"]
		*name = element.(map[string]interface{})["name"]
		*check = element.(map[string]interface{})["is_potentially_hazardous_asteroid"]

		addResponse(date, id, name, check, wg, result)
	}
	mutex.Unlock()
	wg.Done()
}

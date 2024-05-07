package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Response struct {
	Total              interface{} `json:"element_count"`
	Near_earth_objects []struct {
		Date  interface{} `json:"date"`
		ID    interface{} `json:"id"`
		Name  interface{} `json:"name"`
		Check interface{} `json:"is_potentially_hazardous_asteroid"`
	} `json:"near_earth_objects"`
}

const API_KEY = "Th766pSszlQHvMesBM569kkPTJ1UNDsMo6wrLqYV"

var link = fmt.Sprintf("https://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-01&end_date=2024-05-07&api_key=%s", API_KEY)

var result Response

var data map[string]interface{}

func main() {

	// Send GET request
	resp, err := http.Get(link)

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
	go func() { result.Total = data["element_count"] }()

	// Creating a Wait Group
	wg := sync.WaitGroup{}

	// Iteration through data
	for date, value := range data["near_earth_objects"].(map[string]interface{}) {
		var id, name, check interface{}
		wg.Add(1)
		// Concurrent data fetching for each day
		go getDay(&date, &id, &name, &check, value, &wg)
	}
	wg.Wait()

	// Returning a result
	jsonResult, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(string(jsonResult))
}

// Fetching data from one day
func getDay(date *string, id, name, check *interface{}, value interface{}, wg *sync.WaitGroup) {
	for _, element := range value.([]interface{}) {
		*id = element.(map[string]interface{})["id"]
		*name = element.(map[string]interface{})["name"]
		*check = element.(map[string]interface{})["is_potentially_hazardous_asteroid"]

		// Assigning received data to Response object
		result.Near_earth_objects = append(result.Near_earth_objects, struct {
			Date  interface{} "json:\"date\""
			ID    interface{} "json:\"id\""
			Name  interface{} "json:\"name\""
			Check interface{} "json:\"is_potentially_hazardous_asteroid\""
		}{date, id, name, check})
	}
	wg.Done()
}

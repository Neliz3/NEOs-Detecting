package parser

import (
	"NEOs/api"
	"encoding/json"
	"log"
	"sync"
)

type Neo interface {
	AddNeo(obj Near_earth_objects)
}

type Response struct {
	Total              interface{}          `json:"element_count"`
	Near_earth_objects []Near_earth_objects `json:"near_earth_objects"`
}

type Near_earth_objects struct {
	Date  interface{} `json:"date"`
	ID    interface{} `json:"id"`
	Name  interface{} `json:"name"`
	Check interface{} `json:"is_potentially_hazardous_asteroid"`
}

func NewResponse(total interface{}, near_earth_objects []Near_earth_objects) *Response {
	return &Response{
		Total:              total,
		Near_earth_objects: near_earth_objects,
	}
}

func (res *Response) AddNeo(obj Near_earth_objects) []Near_earth_objects {
	res.Near_earth_objects = append(res.Near_earth_objects, obj)
	return res.Near_earth_objects
}

func Parser(body *api.Body) []byte {

	// Decode a Json
	var inputJson map[string]interface{}

	body.Err = json.Unmarshal([]byte(body.Body), &inputJson)
	if body.Err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", body.Err)
	}

	// Assigning a Total value
	response := NewResponse(inputJson["element_count"], []Near_earth_objects{})

	// Creating a Wait Group and Mutex
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	// Iteration through data
	for date, value := range inputJson["near_earth_objects"].(map[string]interface{}) {
		wg.Add(1)

		// Concurrent data fetching for each day
		go getDay(value, &date, &wg, &mutex, response)
	}
	wg.Wait()

	// Returning a result
	jsonResponse, _ := json.MarshalIndent(response, "", "\t")
	return jsonResponse
}

// Fetching data from one day
func getDay(value interface{}, date *string, wg *sync.WaitGroup, mut *sync.Mutex, res *Response) {
	mut.Lock()
	for _, element := range value.([]interface{}) {
		id := element.(map[string]interface{})["id"]
		name := element.(map[string]interface{})["name"]
		check := element.(map[string]interface{})["is_potentially_hazardous_asteroid"]

		neo := Near_earth_objects{Date: date, ID: id, Name: name, Check: check}
		res.AddNeo(neo)
	}
	mut.Unlock()
	wg.Done()
}

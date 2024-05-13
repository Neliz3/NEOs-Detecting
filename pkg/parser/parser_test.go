package parser

import (
	"NEOs/api"
	"io"
	"net/http"
	"testing"

	"github.com/miladibra10/vjson"
)

func Setup(t *testing.T) *api.Body {
	// Read Configs
	link := "https://api.nasa.gov/neo/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&api_key=DEMO_KEY"

	// Send GET request
	resp, err := http.Get(link)

	if err != nil {
		t.Error(err)
	}

	// Read Response
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return &api.Body{Body: bodyResp, Err: err}
}

func TestParser(t *testing.T) {
	json := `{
		"total": 1,
		"near_earth_objects": [
			{
				"date": "2024-05-01",
				"id": "2030825",
				"name": "(1990 TG1)",
				"is_potentially_hazardous_asteroid": false
			}
		]
	}`

	schema, err := vjson.ReadFromString(json)
	if err != nil {
		t.Error(err)
	}

	err = schema.ValidateString(string(Parser(Setup(t))))
	if err != nil {
		t.Error(err)
	}
}

package api

import (
	"NEOs/configs"
	"io"
	"log"
	"net/http"
)

type Body struct {
	Body []byte
	Err  error
}

func GetAPI(link ...string) *http.Response {
	// Read Configs
	configs := configs.ReadConfigs()

	// Send GET request
	if link != nil {
		configs.Link = link[0]
	}
	resp, err := http.Get(configs.Link)

	if err != nil {
		log.Fatalf("No response from request")
	}
	return resp
}

func GetResp() *Body {

	// Read Response
	body, err := io.ReadAll(GetAPI().Body)
	if err != nil {
		log.Fatalf("Unable to read a JSON due to %s", err)
	}

	return &Body{body, err}
}

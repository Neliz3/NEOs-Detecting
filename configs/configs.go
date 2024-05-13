package configs

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	API_KEY string
	Link    string
}

func ReadConfigs(api_key ...string) *Config {
	var API_KEY string

	if api_key != nil {
		API_KEY = api_key[0]
	} else {
		API_KEY = os.Getenv("DEMO_API_KEY")
	}

	config := Config{
		API_KEY: API_KEY,
		Link:    *GenerateLink(&API_KEY),
	}

	return &config
}

func GenerateLink(api_key *string) *string {
	BaseLink := os.Getenv("BaseLink")

	// Time Settings
	today := time.Now().Format("2006-01-02")
	weekAgo := time.Now().AddDate(0, 0, -6).Format("2006-01-02")

	link := fmt.Sprintf(BaseLink, weekAgo, today, *api_key)

	return &link
}

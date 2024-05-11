package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY string
	Link    string
}

func ReadConfigs() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	API_KEY := os.Getenv("API_KEY")

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

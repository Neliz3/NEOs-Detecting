package main

import (
	"NEOs/api"
	"NEOs/pkg/parser"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	fmt.Println(string(parser.Parser(api.GetResp())))

}

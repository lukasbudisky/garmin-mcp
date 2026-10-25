package main

import (
	"log"
	"os"

	"github.com/lukasbudisky/garmin-mcp/garmin"
	"github.com/lukasbudisky/garmin-mcp/mcp"
)

func main() {
	name := "garmin-mcp"
	version := "0.0.1"

	apiKey := os.Getenv("API_KEY")

	token, err := garmin.GetAccessToken(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	err = mcp.StdioServer(&garmin.Auth{Token: token}, name, version, false)
	if err != nil {
		log.Panic(err)
	}

}

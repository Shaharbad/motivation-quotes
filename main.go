package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Quote struct {
	Quote string `json:"quote"`
}

func getRandomQuote(quotes []Quote) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex].Quote
}

func main() {
	// Read quotes from the JSON file
	fileContent, err := ioutil.ReadFile("quotes.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var quotes []Quote
	err = json.Unmarshal(fileContent, &quotes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	app := fiber.New()
	
	app.Get("/quote", func(c *fiber.Ctx) error {
		randomQuote := getRandomQuote(quotes)
		return c.SendString(randomQuote)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3080"
	}
	fmt.Println("Server is listening on 0.0.0.0:" + port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

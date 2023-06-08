package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jugui93/rest-api/database"
	"github.com/jugui93/rest-api/handlers"
	"github.com/jugui93/rest-api/models"
)


func TestListFacts(t *testing.T) {
	// setup
	app := fiber.New()
	handlers := handlers.NewHandlers()
	dsn := fmt.Sprintf(
		"host=dbtest user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME_TEST"),
	)
	database.DB.Connect(dsn)
	SetupRoutes(app, handlers)
	

	// create test data
	fact1 := models.Fact{Question: "What is the capital of France?",
		Answer: "a", A:"Paris", B:"Roma", C:"Barcelona", D:"Munich",Level:1}
	fact2 := models.Fact{Question: "What is the largest ocean?",
		Answer: "d", A:"Artic", B:"Atlantic", C:"Southern", D:"Pacific",Level:1}
	database.DB.Db.Create(&fact1)
	database.DB.Db.Create(&fact2)

	// make request to API
	req := httptest.NewRequest(http.MethodGet, "/fact", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	// check response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	var facts []models.Fact
	if err := json.NewDecoder(resp.Body).Decode(&facts); err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if len(facts) != 3 {
		t.Errorf("Expected 2 facts, but got %d", len(facts))
	}

	if facts[0].Question != fact1.Question || facts[0].Answer != fact1.Answer {
		t.Errorf("Expected fact 1, but got %v", facts[0])
	}

	if facts[1].Question != fact2.Question || facts[1].Answer != fact2.Answer {
		t.Errorf("Expected fact 2, but got %v", facts[1])
	}
}



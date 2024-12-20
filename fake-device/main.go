package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// TemperatureResponse represents the JSON structure for the API response
type TemperatureResponse struct {
	Temperature float64 `json:"temperature"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		// Generate a random temperature between 15.0 and 30.0
		temperature := 15.0 + rand.Float64()*(30.0-15.0)

		// Create the response object
		response := TemperatureResponse{
			Temperature: temperature,
		}

		// Set the Content-Type to application/json
		w.Header().Set("Content-Type", "application/json")

		// Encode the response as JSON and send it
		json.NewEncoder(w).Encode(response)
	})

	// Start the server on port 8080
	http.ListenAndServe(":8085", nil)
}

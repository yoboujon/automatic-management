package server

import (
	"controller/logic"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Payload struct {
	ID    int32 `json:"id"`
	State int32 `json:"state"`
}

func Start(port int64) {
	http.HandleFunc("/sensors", func(w http.ResponseWriter, r *http.Request) {
		handleSensorsRequest(w)
	})
	http.HandleFunc("/actuator", func(w http.ResponseWriter, r *http.Request) {
		handleActuatorRequest(w, r)
	})
	if err := http.ListenAndServe(":"+strconv.FormatInt(port, 10), nil); err != nil {
		panic(err)
	}
}

func handleSensorsRequest(w http.ResponseWriter) {
	var response = logic.GetSensor()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleActuatorRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received ID: %d, State: %v\n", payload.ID, payload.State)
	} else if r.Method == http.MethodGet {
		var response = logic.GetActuator()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

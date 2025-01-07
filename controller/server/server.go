package server

import (
	"controller/logic"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Payload struct {
	State int32 `json:"state"`
}

func Start(port int64) {
	// Creating endpoints
	http.HandleFunc("/sensors", func(w http.ResponseWriter, r *http.Request) {
		handleSensorsRequest(w)
	})
	http.HandleFunc("/actuator/", func(w http.ResponseWriter, r *http.Request) {
		handleActuatorRequest(w, r)
	})

	Logformat(INFO, "Starting server on port %d...\n", port)
	if err := http.ListenAndServe(":"+strconv.FormatInt(port, 10), nil); err != nil {
		Logformat(ERROR, "Could not start server (%s)\n", err.Error())
		panic(err)
	}
}

func handleSensorsRequest(w http.ResponseWriter) {
	Logformat(NOLOG, "GET '/sensors'\n")
	var response = logic.GetSensor()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func actuatorAll(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method == http.MethodPut {
		Logformat(WARNING, "%s '/actuator': (%d) Please provide a specific actuator id.\n", r.Method, http.StatusMethodNotAllowed)
		http.Error(w, "Please provide a specific actuator id", http.StatusMethodNotAllowed)

		//GET
	} else if r.Method == http.MethodGet {
		Logformat(INFO, "%s '/actuator'\n", r.Method)
		var response = logic.GetActuator()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func actuatorSpecific(w http.ResponseWriter, r *http.Request, id int) {
	Logformat(INFO, "%s '/actuator/%d'\n", r.Method, id)

	// POST
	if r.Method == http.MethodPut {
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			Logformat(WARNING, "%s '/actuator': (%d) Could not parse parameters.\n", r.Method, http.StatusBadRequest)
			http.Error(w, "Could not parse parameters", http.StatusBadRequest)
		}
		Logformat(INFO, "%s '/actuator/%d': Changed to %d\n", r.Method, id, payload.State)
		logic.UpdateActuator(id, payload.State)

		// GET
	} else {
		var response = logic.GetActuator()[id]

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleActuatorRequest(w http.ResponseWriter, r *http.Request) {
	if !((strings.Compare(r.Method, http.MethodPut) == 0) || (strings.Compare(r.Method, http.MethodGet) == 0)) {
		Logformat(WARNING, "Invalid method '%s' on '/actuator'\n", r.Method)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	url := strings.Split(r.URL.Path, "/")
	if strings.Compare(url[2], "") == 0 {
		actuatorAll(w, r)
	} else {
		i, err := strconv.Atoi(url[2])
		if err != nil {
			Logformat(WARNING, "%s '/actuator/{id}': You must provide a number (received '%s')\n", r.Method, url[2])
			http.Error(w, "You must provide an id number for the actuator.", http.StatusBadRequest)
		}
		actuatorSpecific(w, r, i)
	}
}

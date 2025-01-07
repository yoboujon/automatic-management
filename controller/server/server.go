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
		handleSensorsRequest(w, r)
	})
	http.HandleFunc("/actuators/", func(w http.ResponseWriter, r *http.Request) {
		handleActuatorRequest(w, r)
	})

	Logformat(INFO, "Starting server on port %d...\n", port)
	if err := http.ListenAndServe(":"+strconv.FormatInt(port, 10), nil); err != nil {
		Logformat(ERROR, "Could not start server (%s)\n", err.Error())
		panic(err)
	}
}

func handleSensorsRequest(w http.ResponseWriter, r *http.Request) {
	if strings.Compare(r.Method, http.MethodGet) != 0 {
		Logformat(WARNING, "Invalid method '%s' on '/sensors'\n", r.Method)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	Logformat(NOLOG, "GET '/sensors'\n")
	var response = logic.GetSensor()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func actuatorAll(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method == http.MethodPut {
		Logformat(WARNING, "%s '/actuators': (%d) Please provide a specific actuator id.\n", r.Method, http.StatusMethodNotAllowed)
		http.Error(w, "Please provide a specific actuator id", http.StatusMethodNotAllowed)
		return

		//GET
	} else if r.Method == http.MethodGet {
		Logformat(INFO, "%s '/actuators'\n", r.Method)
		var response = logic.GetActuators()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			Logformat(WARNING, "%s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func actuatorSpecific(w http.ResponseWriter, r *http.Request, id int) {
	Logformat(INFO, "%s '/actuators/%d'\n", r.Method, id)

	// POST
	if r.Method == http.MethodPut {
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			Logformat(WARNING, "%s '/actuators': (%d) Could not parse parameters.\n", r.Method, http.StatusBadRequest)
			http.Error(w, "Could not parse parameters", http.StatusBadRequest)
			return
		}

		// Updating the actuator if the length is OK
		if logic.UpdateActuator(id, payload.State) {
			Logformat(INFO, "%s '/actuators/%d': Changed to %d\n", r.Method, id, payload.State)
		} else {
			Logformat(WARNING, "%s '/actuators/%d': id too high\n", r.Method, id)
			http.Error(w, "id too high", http.StatusBadRequest)
		}
		return

		// GET
	} else {
		err, response := logic.GetActuator(id)
		if err != nil {
			Logformat(WARNING, "%s '/actuators/%d': %s\n", r.Method, id, err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleActuatorRequest(w http.ResponseWriter, r *http.Request) {
	if !((strings.Compare(r.Method, http.MethodPut) == 0) || (strings.Compare(r.Method, http.MethodGet) == 0)) {
		Logformat(WARNING, "Invalid method '%s' on '/actuators'\n", r.Method)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	url := strings.Split(r.URL.Path, "/")
	if strings.Compare(url[2], "") == 0 {
		actuatorAll(w, r)
	} else {
		i, err := strconv.Atoi(url[2])
		if err != nil {
			Logformat(WARNING, "%s '/actuators/{id}': You must provide a number (received '%s')\n", r.Method, url[2])
			http.Error(w, "You must provide an id number for the actuator.", http.StatusBadRequest)
			return
		}
		actuatorSpecific(w, r, i)
	}
}

package server

import (
	"controller/logic"
	"controller/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func sensorAll(w http.ResponseWriter) {
	util.Logformat(util.NOLOG, "GET '/sensors'\n")
	var response = logic.GetSensors()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sensorSpecific(w http.ResponseWriter, r *http.Request, id int) {
	err, response := logic.GetSensor(id)
	if err != nil {
		util.Logformat(util.WARNING, "%s '/sensors/%d': %s\n", r.Method, id, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.Logformat(util.NOLOG, "GET '/sensors/%d'\n", id)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSensorsRequest(w http.ResponseWriter, r *http.Request) {
	// Checking the method
	if strings.Compare(r.Method, http.MethodGet) != 0 {
		util.Logformat(util.WARNING, "Invalid method '%s' on '/sensors'\n", r.Method)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Checking the state of the URI (/{id} or all)
	check, substr := util.HasSubURI(r)
	if check {
		id, err := strconv.Atoi(substr)
		if err != nil {
			util.Logformat(util.WARNING, "%s '/sensors/{id}': You must provide a number (received '%s')\n", r.Method, substr)
			http.Error(w, "You must provide an id number for the sensor.", http.StatusBadRequest)
			return
		}
		sensorSpecific(w, r, id)
	} else {
		sensorAll(w)
	}
}

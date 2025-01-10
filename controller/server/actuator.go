package server

import (
	"controller/logic"
	"controller/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func actuatorAll(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method == http.MethodPut {
		util.Logformat(util.WARNING, "%s '/actuators': (%d) Please provide a specific actuator id.\n", r.Method, http.StatusMethodNotAllowed)
		http.Error(w, "Please provide a specific actuator id", http.StatusMethodNotAllowed)
		return

		//GET
	} else if r.Method == http.MethodGet {
		util.Logformat(util.NOLOG, "%s '/actuators'\n", r.Method)
		var response = logic.GetActuators()
		sendResponse(w, response)
		return
	}
}

func actuatorSpecific(w http.ResponseWriter, r *http.Request, id int) {
	// PUT
	if r.Method == http.MethodPut {
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			util.Logformat(util.WARNING, "%s '/actuators': (%d) Could not parse parameters.\n", r.Method, http.StatusBadRequest)
			http.Error(w, "Could not parse parameters", http.StatusBadRequest)
			return
		}

		// Updating the actuator if the length is OK
		err, response := logic.UpdateActuator(id, payload.State)
		if err != nil {
			util.Logformat(util.WARNING, "%s '/actuators/%d': id too high\n", r.Method, id)
			http.Error(w, "id too high", http.StatusBadRequest)

		} else {
			util.Logformat(util.NOLOG, "%s '/actuators/%d': state=%d\n", r.Method, id, payload.State)
			util.Logformat(util.INFO, "[%s] Value: %d\n", strings.ToUpper(response.Name), response.Value)
			sendResponse(w, response)
			return
		}
		return

		// GET
	} else {
		util.Logformat(util.NOLOG, "%s '/actuators/%d'\n", r.Method, id)
		err, response := logic.GetActuator(id)
		if err != nil {
			util.Logformat(util.WARNING, "%s '/actuators/%d': %s\n", r.Method, id, err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sendResponse(w, response)
		return
	}
}

func handleActuatorRequest(w http.ResponseWriter, r *http.Request) {
	// Checking the method
	if !((strings.Compare(r.Method, http.MethodPut) == 0) || (strings.Compare(r.Method, http.MethodGet) == 0)) {
		util.Logformat(util.WARNING, "Invalid method '%s' on '/actuators'\n", r.Method)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	check, substr := util.HasSubURI(r)
	if check {
		id, err := strconv.Atoi(substr)
		if err != nil {
			util.Logformat(util.WARNING, "%s '/actuators/{id}': You must provide a number (received '%s')\n", r.Method, substr)
			http.Error(w, "You must provide an id number for the actuator.", http.StatusBadRequest)
			return
		}
		actuatorSpecific(w, r, id)
	} else {
		actuatorAll(w, r)
	}
}

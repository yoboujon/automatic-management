package server

import (
	"controller/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type Payload struct {
	State int32 `json:"state"`
}

func Start(port int64) {
	// Creating endpoints
	http.HandleFunc("/sensors/", func(w http.ResponseWriter, r *http.Request) {
		handleSensorsRequest(w, r)
	})
	http.HandleFunc("/actuators/", func(w http.ResponseWriter, r *http.Request) {
		handleActuatorRequest(w, r)
	})

	util.Logformat(util.INFO, "Starting server on port %d...\n", port)
	if err := http.ListenAndServe(":"+strconv.FormatInt(port, 10), nil); err != nil {
		util.Logformat(util.ERROR, "Could not start server (%s)\n", err.Error())
		panic(err)
	}
}

func sendResponse(w http.ResponseWriter, r any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(r); err != nil {
		util.Logformat(util.ERROR, "%s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

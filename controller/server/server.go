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
	http.HandleFunc("/sensors/", corsMiddleware(handleSensorsRequest))
	http.HandleFunc("/actuators/", corsMiddleware(handleActuatorRequest))

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

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

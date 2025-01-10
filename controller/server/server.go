package server

import (
	"controller/util"
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

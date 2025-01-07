package main

import (
	"controller/logic"
	"controller/server"
)

func main() {
	// Init
	logic.InitSensors()

	// Controller Logic
	logic.StartLogic()

	// Rest API start on port 8085
	server.Start(8085)
}

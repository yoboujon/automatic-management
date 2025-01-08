package main

import (
	"controller/logic"
	"controller/server"
)

func main() {
	// Controller Logic
	logic.StartLogic()

	// Rest API start on port 8085
	server.Start(8085)
}

package main

import (
	"controller/cmd"
	"controller/logic"
	"controller/server"
	"controller/util"
)

func main() {
	// Checking arguments
	args := cmd.CheckArgs()
	util.SetLevel(args.Loglevel)

	// Controller Logic
	logic.StartLogic()

	// Rest API start on port 8085
	server.Start(8085)
}

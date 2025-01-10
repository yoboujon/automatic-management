package cmd

import (
	"controller/util"
	"os"

	"github.com/pborman/getopt/v2"
)

var verboseMap = map[string]util.LogLevel{
	"all":  util.NOLOG,
	"info": util.INFO,
	"warn": util.WARNING,
	"err":  util.ERROR,
}

type Arguments struct {
	Loglevel util.LogLevel
}

func CheckArgs() Arguments {
	optName := getopt.StringLong("verbose", 'v', "info", "default", "all, info, warn, err")
	optHelp := getopt.BoolLong("help", 'h', "Show help")
	getopt.Parse()

	var args = Arguments{}

	// --help/-h
	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	// Checking -v/--verbose
	if len(*optName) != 0 {
		val, ok := verboseMap[*optName]
		if ok {
			util.Logformat(util.INFO, "Logging level set to '%s'.\n", *optName)
			args.Loglevel = val
		} else {
			util.Logformat(util.ERROR, "'%s' is not a valid option for '-v/--verbose'\n", *optName)
			getopt.Usage()
			os.Exit(-1)
		}
	}

	return args
}

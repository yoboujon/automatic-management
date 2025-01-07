package server

import "fmt"

type colorFormat string
type LogLevel int

const (
	BLUE   colorFormat = "\033[34m"
	GREEN  colorFormat = "\033[32m"
	YELLOW colorFormat = "\033[33m"
	RED    colorFormat = "\033[31m"
	RESET  colorFormat = "\033[0m"
)

const (
	INFO LogLevel = iota
	CORRECT
	WARNING
	ERROR
)

var colorPrint = map[LogLevel]colorFormat{
	INFO:    BLUE,
	CORRECT: GREEN,
	WARNING: YELLOW,
	ERROR:   RED,
}

func Logformat(level LogLevel, format string, args ...interface{}) {
	color := colorPrint[level] // Get the color for the given log level
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s%s", color, message, RESET) // Print with color and reset
}

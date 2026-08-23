package main

import (
	"fmt"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

var logLevelNames = map[LogLevel]string{
	Debug:   "DEBUG",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR",
}

func (l LogLevel) String() string {
	return logLevelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Println(level.String())
}

func main() {
	printLogLevel(Debug)
	printLogLevel(Info)
	printLogLevel(Warning)
	printLogLevel(Error)
}

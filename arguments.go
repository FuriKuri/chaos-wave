package main

import (
	"log"
	"os"
	"time"
)

func getArgParameter(name string, defaultValue string) string {
	argsWithoutProg := os.Args[1:]
	for index, element := range argsWithoutProg {
		if element == "--"+name {
			return argsWithoutProg[index+1]
		}
	}
	return defaultValue
}

func interval() time.Duration {
	duration, err := time.ParseDuration(getArgParameter("interval", "10m"))
	if err != nil {
		log.Fatal(err)
	}
	return duration
}

func duration() time.Duration {
	duration, err := time.ParseDuration(getArgParameter("duration", "1h"))
	if err != nil {
		log.Fatal(err)
	}
	return duration
}

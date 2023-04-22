package main

import (
	"os"
)

func main()  {
	// RUN THE PROGRAM USING THIS COMMAND:
	// MONGODB_CONNECTION_STRING='' go run main.go
	if os.Getenv("MONGODB_CONNECTION_STRING") == "" {
		panic("MONGODB_CONNECTION_STRING is not defined")
	}
}

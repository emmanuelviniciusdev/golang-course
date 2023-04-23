package main

import (
	"25mongoapi/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main()  {
	// RUN THE PROGRAM USING THIS COMMAND:
	// MONGODB_CONNECTION_STRING='' go run main.go
	if os.Getenv("MONGODB_CONNECTION_STRING") == "" {
		panic("MONGODB_CONNECTION_STRING is not defined")
	}

	fmt.Println("Welcome to MongoAPI!!")
	fmt.Println("Server running at port 4000")

	log.Fatal(http.ListenAndServe(":4000", routes.Routes()))
}

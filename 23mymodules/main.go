package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	fmt.Println("Hello mod in Golang")

	Greeter()

	router := mux.NewRouter()
	router.HandleFunc("/", ServeHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", router))
}

func Greeter()  {
	fmt.Println("Hey there mod users")
}

func ServeHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to Golang series on YouTube!!</h1>"))
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

const fetchUrl = "https://jsonplaceholder.typicode.com/todos/1"

func main()  {
	fmt.Println("Welcome to web requests in golang!!")

	response, err := http.Get(fetchUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}


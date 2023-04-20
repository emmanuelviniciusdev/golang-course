package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Welcome to web verb video - LCO !!")

	PerformGetRequest()
	PerformPostRequest()
}

func PerformPostRequest()  {
	fmt.Println("Performing POST")

	body := strings.NewReader(`
		{
			"userId": 99,
			"title": "LEARN GOLANG",
			"completed": false
		}
	`)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", body)

	PanicErrIfNotNil(err)

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content-length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	PanicErrIfNotNil(err)

	fmt.Println("Content:", string(content))

	fmt.Println("---")
}

func PerformGetRequest() {
	fmt.Println("Performing GET")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	PanicErrIfNotNil(err)

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content-length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	PanicErrIfNotNil(err)

	fmt.Println("Content:", string(content))

	fmt.Println("---")
}

func PanicErrIfNotNil(err any) {
	if err != nil {
		panic(err)
	}
}

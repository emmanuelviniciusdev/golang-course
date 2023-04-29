package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

var mutex sync.Mutex

var reachedEndpoints []string

func main() {
	websiteList := []string{
		"https://icloud.com",
		"https://emmanuelbergmann.me",
		"https://github.com",
		"https://invalid.invalid",
		"https://fb.com",	
		"aaaaa this is super invalid",
	}

	for _, website := range websiteList {
		go getStatusCode(website, &mutex)
		waitGroup.Add(1)
	}

	waitGroup.Wait()

	fmt.Println("reachedEndpoints:", reachedEndpoints)
}

func getStatusCode(endpoint string, mutex *sync.Mutex) {
	defer waitGroup.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Something went wrong trying to reach '%s'\n", endpoint)
		return
	}

	fmt.Printf("Status code for '%s' is '%d'\n", endpoint, res.StatusCode)
	
	mutex.Lock()
	reachedEndpoints = append(reachedEndpoints, endpoint)
	mutex.Unlock()
}

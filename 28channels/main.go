package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	
	myChannel := make(chan int)

	wg.Add(2)

	go func (wg *sync.WaitGroup, ch chan int) {
		channelValue, isChannelOpened := <- ch

		if isChannelOpened {
			fmt.Println("Channel is opened and its value is:", channelValue)
		} else {
			fmt.Println("Channel is closed")
		}


		wg.Done()
	}(&wg, myChannel)

	go func (wg *sync.WaitGroup, ch chan int) {
		ch <- 0
		close(ch)

		wg.Done()
	}(&wg, myChannel)

	wg.Wait()
}

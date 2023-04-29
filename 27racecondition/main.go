package main

import (
	"fmt"
	"sync"
)

func main () {
	var wg sync.WaitGroup
	var m sync.Mutex

	score := []int{0}

	wg.Add(3)

	go func (wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Race 1")

		m.Lock()
		score = append(score, 1)
		m.Unlock()

		wg.Done()
	}(&wg, &m)
	
	go func (wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Race 2")

		m.Lock()
		score = append(score, 2)
		m.Unlock()

		wg.Done()
	}(&wg, &m)

	go func (wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Race 3")

		m.Lock()
		score = append(score, 3)
		m.Unlock()

		wg.Done()
	}(&wg, &m)

	wg.Wait()

	fmt.Println("Score:", score)
}

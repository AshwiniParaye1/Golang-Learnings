package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mutexx := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mutexx.Lock()
		score = append(score, 1)
		mutexx.Unlock()
		wg.Done()
	}(wg, mutexx)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mutexx.Lock()
		score = append(score, 2)
		mutexx.Unlock()
		wg.Done()

	}(wg, mutexx)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mutexx.Lock()
		score = append(score, 3)
		mutexx.Unlock()
		wg.Done()

	}(wg, mutexx)

	wg.Wait()
	fmt.Println(score)

}

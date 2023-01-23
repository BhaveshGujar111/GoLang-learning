package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - lco.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	// mut.RLock()
	var score = []int{0}
	// mut.RUnlock()

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One R")

		mut.RLock()
		score = append(score, 1)
		mut.RUnlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two R")

		mut.RLock()
		score = append(score, 2)
		mut.RUnlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three R")

		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three R")

		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()

}

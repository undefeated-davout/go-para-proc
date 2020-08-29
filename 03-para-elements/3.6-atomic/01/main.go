package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup
	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		atomic.AddInt64(&count, 1)
		fmt.Printf("Incrementing: %d\n", count)
	}

	wg.Add(1)
	go increment(&wg)
	wg.Wait()
}

func mainAnother() {
	var count int64
	var lock sync.Mutex
	var wg sync.WaitGroup
	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	wg.Add(1)
	go increment(&wg)
	wg.Wait()
}

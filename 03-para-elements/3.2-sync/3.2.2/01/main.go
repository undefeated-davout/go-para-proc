package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// インクリメント
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// デクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Printf("count: %d\n", count)
	fmt.Println("Arithmetic complete.")
}

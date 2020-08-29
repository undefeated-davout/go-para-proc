package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once
	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment) // 異なるゴルーチンで呼び出されても1回だけ実行される
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count) // Count is 1
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	// わざわざポインタ引数でWaitGroupを渡す必要ある？
	// 情報のやりとりが分かりやすくなって良いかも？
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters) // まとめてゴルーチンを監視
	for i := 0; i < numGreeters; i++ {
		// wg.Add(1) とforの中に書いても良い
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func mainAnother() {
	const numGreeters = 5
	var wg sync.WaitGroup
	hello := func(id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	wg.Add(numGreeters) // まとめてゴルーチンを監視
	for i := 0; i < numGreeters; i++ {
		go hello(i + 1)
	}
	wg.Wait()
}

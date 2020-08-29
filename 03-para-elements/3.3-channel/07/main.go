package main

import (
	"fmt"
	"time"
)

func main() {
	// resultStreamのライフサイクルはこの関数内で完結する
	chanOwner := func() <-chan int { // 読み込み専用のチャネルを返す
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream) // チャネルの所有者なので責任持ってclose
			for i := 0; i <= 5; i++ {
				fmt.Printf("resultStream <- i 直前: %d\n", i)
				resultStream <- i // 6つ目の値がブロックされ、読み込み(range)時点で値が入る
				fmt.Printf("resultStream <- i 直後: %d\n", i)
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	time.Sleep(1 * time.Second)
	fmt.Println("range resultStream 開始")
	for result := range resultStream { // チャネルを読み込むだけ
		time.Sleep(1 * time.Second)
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

// resultStream <- i 直前: 0
// resultStream <- i 直後: 0
// resultStream <- i 直前: 1
// resultStream <- i 直後: 1
// resultStream <- i 直前: 2
// resultStream <- i 直後: 2
// resultStream <- i 直前: 3
// resultStream <- i 直後: 3
// resultStream <- i 直前: 4
// resultStream <- i 直後: 4
// resultStream <- i 直前: 5
// range resultStream 開始
// resultStream <- i 直後: 5
// Received: 0
// Received: 1
// Received: 2
// Received: 3
// Received: 4
// Received: 5
// Done receiving!

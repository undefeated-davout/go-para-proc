package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 1つのゴルーチンが起動したことを表す
	// ゴルーチンの外で行うことに注意
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // 全てのゴルーチンが終了するまでメインゴルーチンをブロック
	fmt.Println("All goroutines complete.")
}

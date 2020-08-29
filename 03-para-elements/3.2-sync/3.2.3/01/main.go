package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	// 長さ0のスライス（10回足すのでキャパシティ10）
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // 条件のクリティカルセクションに入る
		queue = queue[1:] // 2番目以降をセット（1番目を消す）
		fmt.Println("Remove from queue")
		c.L.Unlock() // 要素をキューから取り出したのでクリティカルセクションを抜ける
		c.Signal()   // 条件を待っているゴルーチンに何かが起きたことを知らせる
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // クリティカルセクションに入る
		for len(queue) == 2 {
			c.Wait() // 条件のシグナルが出るまでメインゴルーチンを停止
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock() // クリティカルセクションを抜ける
	}
}

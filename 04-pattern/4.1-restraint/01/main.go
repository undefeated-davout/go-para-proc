package main

import "fmt"

func main() {
	// 読み込み専用チャネルを戻り値に取ることで
	// スコープ外でのチャネルへの書き込みを防ぐ
	chanOwner := func() <-chan int {
		// チャネルの書き込み権限を拘束する
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// 読み込み専用チャネルを引数に取ることで
	// 操作を読み込み専用に拘束する
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}

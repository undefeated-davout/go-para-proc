package main

import "fmt"

func main() {
	// valueStream := make(chan interface{})
	// close(valueStream)

	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream // 閉じたチャネルから読み込むことができる
	fmt.Printf("(%v): %v", ok, integer)
}

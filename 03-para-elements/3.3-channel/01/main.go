package main

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		stringStream <- "Hello"
	}()
	fmt.Println(<-stringStream) // ブロックされる
}

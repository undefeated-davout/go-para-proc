package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		stringStream <- "Hello"
	}()
	fmt.Println(<-stringStream) // stringStreamに値がセットされることはないのでデッドロック
}

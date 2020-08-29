package main

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		stringStream <- "Hello"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}

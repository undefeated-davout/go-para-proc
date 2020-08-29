package main

import (
	"fmt"
	"time"
)

func main() {
	var data int

	go func() {
		data++
	}()
	time.Sleep(1 * time.Second) // 間違った解決策
	fmt.Printf("the value is %v.\n", data)
}

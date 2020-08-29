package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// intStreamがcloseされたことを持ってrangeが開始される
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

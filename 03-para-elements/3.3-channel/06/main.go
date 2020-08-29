package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			<-begin // この行でclose(begin)されるまでブロック
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin) // <-beginにシグナルを送る
	wg.Wait()
}

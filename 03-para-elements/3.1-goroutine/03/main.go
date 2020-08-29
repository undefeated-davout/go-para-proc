package main

import (
	"fmt"
	"sync"
)

func mainNg() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// [good day]が3回表示される（rangeが全部回り切ってから実行される）
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func mainOk() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, salutation := range []string{"hello", "greetings", "good day"} {
			fmt.Println(salutation)
		}
	}()
	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // funcの引数に変数を渡すと、その時点の変数の値で動作する
	}
	wg.Wait()
}

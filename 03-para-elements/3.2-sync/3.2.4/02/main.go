package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment) // 最初のDoだけ実行される
	once.Do(decrement) // 実行されない

	fmt.Printf("Count: %d\n", count)
}

func sample() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	// このDoは、下の[onceA.Do(initA)]が終了するまで実行できない（同じonceAだから）
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
	// このプログラムはデッドロックする
}

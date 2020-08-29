package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // 1つめのインスタンス(Create)
	instance := myPool.Get() // 2つめのインスタンスを生成(Create)
	myPool.Put(instance)     // インスタンスをプールに戻す
	myPool.Get()             // Createされない。先に生成されてプールに戻されたインスタンスを再利用
}

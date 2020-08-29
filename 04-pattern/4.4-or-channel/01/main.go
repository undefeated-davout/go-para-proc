package main

import (
	"fmt"
	"time"
)

func main() {
	// 引数：チャネルの可変長引数のスライス
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
					// スライスの3番目以降のチャネルから再起的にorチャネル作成
					//スライスの残り部分をorチャネルに分解して、最初のシグナルが返る木構造
					// orDoneチャネルも渡して、木構造の上位部分が終了したら下位部分も終了
				}
			}
		}()
		return orDone
	}

	// 指定秒経過したら終了（読み取りチャネルを返す）
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// どれかがチャネルを返したら終了（最速のやつ）
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

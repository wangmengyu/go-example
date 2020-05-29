package main

import (
	"fmt"
	"time"
)

/**
  timers 是用于将来某个时刻执行一次.
  tickers是你想要做某些事情重复的 周期性的
*/
func main() {
	//ticker 使用一个类似 timers的机制,
	// 一个channel 被发送值.
	//这里我们将使用select 通道来等待, 没500ms一次
	tickers := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) // 通知完成的通道
	//不断的从tickers的通道获取消息.直到收到结束通知为止
	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-tickers.C:
				fmt.Println("Tick at:", v)
			}
		}
	}()

	//Tickers 可以被stop, 一旦一个ticker被stop, 它就不会再收到时间消息, 我们将在1600ms后stop
	time.Sleep(time.Millisecond * 1600)
	done <- true
	tickers.Stop()
	fmt.Println("stop tickers")

}

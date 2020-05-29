package main

import (
	"fmt"
	"time"
)

/**
  限制频率 是一个很重要的机制, 对于控制资源, 和维护服务质量,
  go 支持优雅的限制频率 在goroutie , channel 和 tickers
*/
func main() {
	//首先先收集请求到channel中, 关闭通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	//limiter channel 每隔200毫秒 收到一个时间
	limiter := time.Tick(200 * time.Millisecond)

	//为了限制频率, 每次从limiter接收到事件, 才处理接收到的数据
	for v := range requests {
		t := <-limiter
		fmt.Printf("get val %d from request time:%v\n", v, t)
	}

	//我们可能想要允许请求中的小爆发,
	//同时保留总体的速率限制, 我们可以用缓冲限制器来完成这个
	burstyLimiter := make(chan time.Time, 3)
	//填充数据到缓存中
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每200毫秒, 我们变价一个新的元素到bustryLimit里去
	//并发的
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//现在 完成5个输入请求,
	bustryRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bustryRequests <- i
	}

	close(bustryRequests)
	for n := range bustryRequests {
		t := <-burstyLimiter
		fmt.Printf("request = %d, time=%v\n", n, t)
	}
}

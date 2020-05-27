package main

import (
	"fmt"
	"time"
)

func main() {
	// channel 同步机制

	// 创建一个done的channel , 用来通知任务完成
	done := make(chan bool)
	// 创建一个异步的worker , 完成指定任务,完成时通知给刚才创建的通道
	go func() {
		fmt.Println("working...")
		time.Sleep(1 * time.Second)
		fmt.Println("done")
		done <- true
	}()
	<-done // 阻塞, 直到有信号收到才执行下去. 去掉这一行 goroutine里面的就不执行了.
	fmt.Println("after done")

}

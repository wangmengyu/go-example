package main

import (
	"fmt"
	"time"
)

/**
  我们经常执行GO代码在将来的某个时间执行,
  或者在一段时间内重复执行,
  Go内置的 timer 和 ticker 特色 都可以让完成这些任务
  我们将先看一下timers
*/
func main() {
	//Timers 表示一个独立的事件 在将来, 你告诉timer你要等待多久,
	// 还有 他提供一个channel 将在指定的时间内被通知到.
	timer1 := time.NewTimer(2 * time.Second) // 提供一个管道, 2秒后收到通知

	//<=timer1.C 在 timers的管道C 上 阻塞, 等到两秒后 会对C发送一个值指示, 标志着timer到期了
	<-timer1.C

	fmt.Println("Timer 1 fired") // 过期了

	//如果你指示想要等待, 你可以使用time.Sleep ,
	// 例如 想要在过期之前, 先取消掉timer .
	timer2 := time.NewTimer(1 * time.Second)
	//并发的, 如果1秒到了. 它会取消
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	//中断timer 2
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	//给timer2足够时间来过期
	//time.Sleep(2*time.Second)

}

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/**
  使用goroutine和通道内置同步特性来实现同步跨多个goroutine对共享状态的访问
  这种通过channel的方法与GO的共享内存的思想一致,
  通过交流共享内存
*/

/**
我们将状态归属于一个独立的goroutine,
这将保证数据不会被并行访问所覆盖
为了读写状态,
其他goroutine将发送数据给拥有者状态的goroutine,并且接收相应的答复,
这些readOp和writeOp结构封装了这些请求
并且为拥有者goroutine 提供相应
*/

type readOp struct { // 通知读请求的操作
	key  int
	resp chan int //响应  , 通知需要读取的key是什么
}

type writeOp struct { //通知写请求的操作
	key  int
	val  int
	resp chan bool //响应 , 通知完成了写操作
}

func main() {

	//计数执行了多少操作
	var readOps uint64
	var writeOps uint64

	// 读和写用独立的通道, 用于发出读和写的请求
	readChannel := make(chan readOp)
	writeChannel := make(chan writeOp)

	//下面这个保存状态的goroutine,
	// 不断的接收读和写的请求, 并且有的时候把消息推送到resp

	var state = make(map[int]int)
	go func(map[int]int) {
		//var state = make(map[int]int)
		for {
			select {
			case read := <-readChannel: //读取到的相关KEY的在MAP中的数据, 写入到read.resp中
				read.resp <- state[read.key]
			case write := <-writeChannel: //写入的新数据, 写入到state中, 并且告知write.resp完成了操作
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}(state)

	//开启100个并发, 每个并发不断的来给读通道发送请求, 每个请求有1毫秒的延时
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				readChannel <- read // 发送好之后, 阻塞等待处理完成
				<-read.resp
				//原子计算读的次数
				atomic.AddUint64(&readOps, 1)
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	//开10个并发, 每个并发不断的创建写请求, 并且接收写入完成状态,
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writeChannel <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	//打印读的次数,写次数, 当前Map
	readFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read final:", readFinal)
	writeFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write final:", writeFinal)
	fmt.Println("state:", state)

}

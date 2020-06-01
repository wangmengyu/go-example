package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	state := make(map[int]int)
	var writeNum uint64
	var readNum uint64

	readCh := make(chan readOp)
	writeCh := make(chan writeOp)

	/**
	  不断从读和写的channel获取要处理的数据
	*/
	go func(map[int]int) {
		for {
			select {
			case read := <-readCh:
				read.resp <- state[read.key]
				atomic.AddUint64(&readNum, 1)

			case write := <-writeCh:
				//write val to state
				state[write.key] = write.val
				write.resp <- true
				atomic.AddUint64(&writeNum, 1)

			}
		}
	}(state)

	//开100个读并发,每个并发不断的往readCh放入数据
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				readCh <- read
				<-read.resp
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	//开10个并发, 每个并发不断的往writeCh发送数据
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writeCh <- write
				<-write.resp
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	time.Sleep(1 * time.Second)

	//打印读的次数
	fmt.Println("read num: ", atomic.LoadUint64(&readNum))
	fmt.Println("write num: ", atomic.LoadUint64(&writeNum))
	fmt.Println("state:", state)

}

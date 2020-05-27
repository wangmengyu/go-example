package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	//并发的向两个通道同时发送数据
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			c1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			c2 <- "channel 2"
		}
	}()

	//select 可以用来等待多个通道的数据,
	//通常, select 外面要套一个循环. 循环次数= 所有数据发送次数的综合, 这里是两次
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

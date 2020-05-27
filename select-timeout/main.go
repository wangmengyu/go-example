package main

import (
	"fmt"
	"time"
)

func main() {

	//time.After可以用于超时处理.
	//创建一个通道, 2秒后对他输入一个数据
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "channel 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	//创建一个(通道2, 1秒后对他进行输入一个数据, 3秒内超时
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "channel 2"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

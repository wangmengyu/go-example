package main

import "fmt"

/**
  用于接收消息
*/
func ping(pings chan<- string, msg string) {
	pings <- msg
}

/**
  用于吧消息从接收管道发送到响应管道
*/
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed msg")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

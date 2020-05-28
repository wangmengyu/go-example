package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// send data to a no buff channel , can not send ,
	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("send string", msg)
	default:
		fmt.Println("no send msg")
	}

	// get data from messages channel
	select {
	case msg := <-messages:
		fmt.Println("get data from messages", msg)
	case sig := <-signals: //  取不到也不会报deadlock
		fmt.Println("get data from signal", sig)
	default:
		fmt.Println("no activity")
	}

}

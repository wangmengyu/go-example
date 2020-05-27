package main

import "fmt"

func main() {
	// create a string channel
	message := make(chan string)
	//并发的 put data into channel
	go func() {
		message <- "ping"
	}()

	//receive data from message
	msg := <-message
	fmt.Println(msg)

}

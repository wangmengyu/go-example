package main

import "fmt"

func main() {
	//create a channel - with 2
	c := make(chan string)
	go func() {
		c <- "hi"
		c <- "hello"
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)

}

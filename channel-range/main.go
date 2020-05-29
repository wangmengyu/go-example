package main

import "fmt"

func main() {

	messages := make(chan string)

	messages <- "one"
	messages <- "two"

	// 关闭的通道还是可以取出数据的
	close(messages)
	for v := range messages {
		fmt.Println(v)
	}

}

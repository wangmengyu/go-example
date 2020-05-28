package main

import "fmt"

func main() {

	jobs := make(chan int)
	done := make(chan bool)
	//开个goroutine, 从jobs通道获取数据, 如果通道关闭了, 通知done通道
	go func() {
		for {
			if val, ok := <-jobs; ok { // 这里的ok, 在jobs被close后会变成false, 否则一直是true
				fmt.Println("val = ", val)
			} else {
				done <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs)

	<-done // 没有信号告知结束是不能从这里输出数据的.在这之前必须收的到信号
	fmt.Println("done")

}

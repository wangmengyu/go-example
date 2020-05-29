package main

import (
	"fmt"
	"time"
)

/**
  如何实现一个 worker pool ,
  使用goroutine & channel
*/

/**
  这里的worker方法.
  我们将会执行几个并发的实例 (外面用go调用)
  这些workers将从jobs通道接收任务, 然后 发送 相应的结果到results通道,
  我们会sleep一秒 在每个job执行过程中, 来效仿完成任务
*/

func worker(id int, in <-chan int, out chan<- int) {
	for v := range in {
		fmt.Printf("workder %d  receive %d\n", id, v)
		time.Sleep(1 * time.Second)
		out <- v * 2
		fmt.Printf("worker %d output %d\n", id, v*2)
	}
}

func main() {
	//为了使用我们的worker池, 我们需要发送数据给他们工作, 还有收集他们处理的结果
	//定义worker的数量
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	s1 := time.Now().Unix()

	//开3个worker一起完成处理
	for j := 1; j <= 3; j++ {
		go worker(j, jobs, results)
	}

	//输入要处理的数据
	for k := 1; k <= numJobs; k++ {
		jobs <- k
	}

	//结果的输出
	for p := 1; p <= numJobs; p++ {
		n := <-results
		fmt.Println("result:", n)
	}

	s2 := time.Now().Unix()
	fmt.Println("s2-s1=", s2-s1)

}

package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  为了等待多个goroutine完成, 我们可以使用wait group
*/

//这个方法里, 是在每一个goroutine里的事情
//注意, 一个waitgroup必须用指针形式传入

func worker(id int, wg *sync.WaitGroup) {
	//返回之前, 要通知waitgroup已经完成了
	defer wg.Done()

	fmt.Printf("Worder %d starting\n", id)

	//sleep一秒模仿完成了任务
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {

	//这个 waitgroup 是用于瞪大所有的goroutine运行完成.
	var wg sync.WaitGroup

	//分配任务,每配一个需要增加wg的计数
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	//阻塞, 知道所有任务完成通知wg
	wg.Wait()

}

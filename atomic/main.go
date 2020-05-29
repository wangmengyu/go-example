package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
 在GO中, 主要的管理状态的机制是通过channel来交流信息,
 我们看到worker-pool的范例,
还有 一些其他得到选项来管理状态,
这里我们看一下使用sync/atomic包,
atomic counters 处理多个goroutine
*/
func main() {

	// 我们将使用一个无符号整数来表达我们的计数器
	var ops uint64

	//waitgroup会帮助我们等待所有的goroutine完成他们的工作
	var wg sync.WaitGroup

	//我们会开始 50个 goroutine , 每个增加计数器 1000 次

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				//为了原子的叠加计数器, 我们要使用AddUnit64, 要给地址,
				atomic.AddUint64(&ops, 1)
				v := atomic.LoadUint64(&ops)
				fmt.Println("v=", v)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// 这里访问ops是安全的因为我们知道没有其他goroutine等待他, 如果需要, 可以用读取原子安全可以用atomic.LoadUint64
	fmt.Println(ops)

}

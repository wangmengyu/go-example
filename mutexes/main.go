package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
  控制单个计数器的原子性可以用atomic
  对更加复杂的状态, 我们可以使用 mutex 安全的访问数据 通过多个goroutines
*/
func main() {
	// 这里用一个map来演示
	state := make(map[int]int)
	// mu 将同步的访问state
	mu := &sync.Mutex{}

	// 我们将跟踪多少次读写
	var readOps uint64
	var writeOps uint64

	//我们开始100个goroutines 来执行重复读取, 每1000毫秒在每个goroutine中
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				//随机一个key [1-5]
				key := rand.Intn(5)
				//锁定
				mu.Lock()
				total += state[key]
				mu.Unlock()
				atomic.AddUint64(&readOps, 1)
				//每次读取间隔1毫秒
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//同时 开启10个goroutine 进行写操作,
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mu.Lock()
				state[key] = val
				mu.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(1 * time.Second)
	//跑1分钟
	//打印当前的readOps
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read ops:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write ops:", writeOpsFinal)

	//打印state
	mu.Lock()
	fmt.Println("state:", state)
	mu.Unlock()

}

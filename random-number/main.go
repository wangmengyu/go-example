package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64()) // rand float val from 0 to 1

	fmt.Println(rand.Float64()*5 + 5)

	//默认的数字生成器是固定的,每次默认情况会生成相同的数字序列
	//要产生不同序列, 可以给他一个可以改变的种子,
	//对于大蒜保密的随机数字, 使用这种方法是不安全的,需要对数字使用加密
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))

	//用常量生成种子
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))

	//使用相同的数字作为种子, 他会产生相同的随机数字序列
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100))
	fmt.Println(r3.Intn(100))
}

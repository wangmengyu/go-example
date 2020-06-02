package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano() //unix纳秒
	fmt.Println(now)        // 自unix时代来经过的时间

	millis := nanos / 1000000 //
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
	s1 := rand.NewSource(nanos)
	fmt.Println(rand.New(s1))

}

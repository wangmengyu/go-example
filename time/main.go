package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)
	p(now)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now)) // 是否在当前时间之前
	p(then.After(now))  // 是否在当前时间之前
	p(then.Equal(now))  // 是否在当前时间之前

	diff := now.Sub(then)
	p(diff)         // 相差时间
	p(diff.Hours()) // 相差小时
	p(diff.Minutes())
	p(diff.Seconds())

	p(then.Add(diff)) // 当前时间
	p(then.Add(-diff))

}

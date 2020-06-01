package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort类对内建类型的排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs) // 直接修改原来的slice

	ints := []int{7, 2, 4}

	sort.Ints(ints)
	fmt.Println("ints:", ints)

	//检查是不是已经排好序的
	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}

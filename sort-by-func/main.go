package main

import (
	"fmt"
	"sort"
)

//自定义排序需要自定义类型来进行排序
type byLength []string

//实现自定义排序., 需要实现Len, Swap, Less3个接口
//这样就能够使用Sort的泛型函数
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//长度小的认为是小得
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	//自定义排序规则
	//例如:根据字符长度来进行排序
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))

	fmt.Println(fruits)

}

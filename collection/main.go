package main

import (
	"fmt"
	"strings"
)

//查询, 从vs中找出t的节点位置
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

/**
是否包含
*/
func Include(vs []string, t string) bool {
	return Index(vs, t) > 0
}

/**
  只要有任何一个元素符合f函数判定, 马上返回true
*/
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

/**
所有元素都符合某个函数
*/

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

/**
返回新的片, 其中包含符合指定规则的元素
*/
func Filter(vs []string, f func(string) bool) []string {
	res := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

/**
根据指定的方法转换每个元素到新的片
*/
func Map(vs []string, f func(string) string) []string {
	res := make([]string, 0)
	for _, v := range vs {
		res = append(res, f(v))
	}
	return res
}
func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "plum"))

	//是否有任何一个p开头的元素
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	//是否全部是p开头的
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	//值保留含有e的元素
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))

	//把所有元素变成大写
	fmt.Println(Map(strs, strings.ToUpper))

}

package main

/**
  在这个范例中我们的state会属于独立的goroutine
  他会确保数据不会被覆盖从并发的请求中,
  为了读写state, 其他goroutine会发送消息到自己的goroutine ,
*/
func main() {

}

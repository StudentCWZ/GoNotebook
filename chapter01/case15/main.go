/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 10:57
*/

package main

/*
	1. 通道(channel)和 goroutine 搭配，实现用通信代替内存共享的 CSP 模型。
*/

// 消费者
func consumer(data chan int, done chan bool) {
	for x := range data { // 接收数据，直到通道关闭
		println("recv:", x)
	}

	done <- true // 通知 main，消费结束
}

// 生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i // 发送数据
	}
	close(data) // 生产结束，关闭通道
}

func main() {
	done := make(chan bool) // 用于接收消费结束信号
	data := make(chan int)  // 数据管道

	go consumer(data, done) // 启动消费者
	go producer(data)       // 启动生产者

	<-done // 阻塞，直到消费者发回信结束信号
}

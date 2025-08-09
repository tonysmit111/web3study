package main

import (
	"fmt"
	"time"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func producer(ch chan int) {
	go func() {
		for i:=1;i<=100;i++ {
			ch <- i
			if i%10 == 0 {
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}

func consumer(ch chan int) {
	go func() {
		for v:= range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	ch:=make(chan int,5)
	producer(ch)
	consumer(ch)
	time.Sleep(time.Second * 10)
}
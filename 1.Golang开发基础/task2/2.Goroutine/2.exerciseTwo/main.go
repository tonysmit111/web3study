package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

func executor(tasks *[]func()) {
	var wg sync.WaitGroup
	wg.Add(len(*tasks))
	begin:=time.Now().UnixMilli()
	for k, v := range *tasks {
		go func() {
			start := time.Now().UnixMilli()
			v()
			fmt.Println(k, "号任务执行时间：", time.Now().UnixMilli()-start)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("所有任务总体执行时间：",time.Now().UnixMilli()-begin)
}

func main() {
	executor(&[]func(){func() { time.Sleep(time.Second * 5) }, func() { time.Sleep(time.Second * 3) }})
}

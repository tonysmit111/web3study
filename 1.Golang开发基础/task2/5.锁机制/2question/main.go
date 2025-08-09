package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

var wg sync.WaitGroup

type Counter struct {
	count int64
}

func (c *Counter) increase() {
	defer wg.Done()
	for range 1000 {
		atomic.AddInt64(&(c.count), 1)
	}
}

func main() {
	c := &Counter{}
	for range 10 {
		wg.Add(1)
		go c.increase()
	}
	wg.Wait()
	fmt.Println(c.count)
}

package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

var wg sync.WaitGroup
var lock sync.Mutex

func add(num *int){
	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()
	for range 1000 {
		*num += 1
	}
}


func main() {
	n := 0
	for i:=0;i<10;i++ {
		wg.Add(1)
		go add(&n)
	}
	wg.Wait()
	fmt.Println(n)
}



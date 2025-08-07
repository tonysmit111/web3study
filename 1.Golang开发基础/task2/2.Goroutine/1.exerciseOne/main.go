package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

var wg sync.WaitGroup

func printOdd(n int) {
	for i := 1; i <= n; i += 2 {
		fmt.Println(i)
	}
	wg.Done()
}

func printEven(n int) {
	for i := 2; i <= n; i += 2 {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go printOdd(10)
	go printEven(10)
	wg.Wait()
}

package main

import "fmt"

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
// func multiplyTwo(arr *[]int) {
// 	arr2 := *arr
// 	for i:=0;i<len(arr2);i++ {
// 		arr2[i] *= 2
// 	}
// }

func multiplyTwo(arr *[]int) {
	for i:=0;i<len(*arr);i++ {
		(*arr)[i] *= 2
	}
}

func main() {
	arr := []int{1,2,3}
	multiplyTwo(&arr)
	fmt.Println(arr)
}


package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
}

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
func singleNumber(nums []int) int {
    m := make(map[int]int, 0)
	for _, i := range nums {
		m[i]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	panic("no single number found")
}

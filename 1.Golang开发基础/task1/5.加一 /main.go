package main

import "fmt"

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
// 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func main() {
	arr := []int{9, 9, 9, 9, 9, 9}
	fmt.Println(plusOne(arr))
}

func plusOne(digits []int) []int {
	tmp := reverse(digits)
	x := false
	for i := 0; i < len(tmp); i++ {
		b := tmp[i]+1 > 9
		if b {
			tmp[i] = 0
		} else {
			x = true
			tmp[i] = tmp[i] + 1
			break
		}
	}
	if !x {
		tmp = append(tmp, 1)
	}
	return reverse(tmp)
}

func reverse(digits []int) []int {
	rev := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		rev = append(rev, digits[i])
	}
	return rev
}

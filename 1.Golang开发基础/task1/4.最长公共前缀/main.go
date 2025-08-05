package main

import (
	"fmt"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func main() {
	// fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	// fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	// fmt.Println(longestCommonPrefix([]string{"reflower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"a","b"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first:=strs[0]
	if len(strs) == 1 {
		return first
	}
	prefix := first
	for _,v:=range strs {
		if len(v) < len(prefix) {
			prefix = v
		}
	}
	for i := 0; i < len(strs); i++ {
		for j := 1; j <= len(prefix); j++ {
			b := strs[i][:j] == prefix[:j]
			if !b {
				prefix = strs[i][:j-1]
				break
			}
		}
	}
	return prefix
}

package main

import (
	"fmt"
	"math"
	"sort"
)

type ST [][]int

func (st ST) Len() int{
	return len(st)
}

func (st ST) Swap(i, j int) {
	st[i],st[j] = st[j], st[i]
}

func (st ST) Less(i, j int) bool{
	return st[i][0] < st[j][0]
}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	sort.Sort(ST(intervals))
	slice := [][]int{}
	for i := 0; i < len(intervals); i++ {
		lastIndex := len(slice) -1
		processInter := intervals[i]
		if lastIndex <= -1 {
			slice = append(slice, processInter)
			continue
		}
		end := slice[lastIndex][1]
		processStart := processInter[0]
		if processStart > end  {
			slice = append(slice, processInter)
		} else {
			processEnd := processInter[1]
			newEnd := math.Max(float64(end), float64(processEnd))
			slice[lastIndex][1] = int(newEnd)
		}
	}
	return slice
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	sort.Sort(ST(intervals))
	fmt.Println(intervals)
	fmt.Println(merge(intervals))
}

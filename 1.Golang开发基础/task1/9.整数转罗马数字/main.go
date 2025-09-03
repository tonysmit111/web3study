package main

import (
	"fmt"
	"strings"
)

type KV struct  {
	K int
	V string
}

// 如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
// 如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号，例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX。
// 仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)。
/*
func intToRoman(num int) string {
    // m := make(map[int]string)
	// m[1]="I"
	// m[5]="V"
	// m[10]="X"
	// m[50]="L"
	m := map[int]string{1000:"M",900:"CM",500:"D",400:"CD",100:"C",90:"XC",50:"L",40:"XL",10:"X",9:"IX",5:"V",4:"IV",1:"I"}
	ks := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}

	var strSlice []string
	for _,k := range ks {
		for ;num >= k; {
			num = num - k
			strSlice = append(strSlice, m[k])
		}
	}
	fmt.Println(strSlice)
	return strings.Join(strSlice, "")
}
*/
func intToRoman(num int) string {
    m := map[int]string{1000:"M",900:"CM",500:"D",400:"CD",100:"C",90:"XC",50:"L",40:"XL",10:"X",9:"IX",5:"V",4:"IV",1:"I"}
	ks := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}

	var str string
	for _,k := range ks {
		for ;num >= k; {
			num = num - k
			str = str + m[k]
		}
	}
	return str
}


func main() {
	s := intToRoman(3494)
	fmt.Println((s))
}
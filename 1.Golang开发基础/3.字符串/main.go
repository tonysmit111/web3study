package main

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {
    m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	sli := []string{}
	for _, v := range s {
		rightStr := m[string(v)]
		if rightStr != "" {
			sli = append(sli, string(v))
		} else {
            if len(sli) <= 0 {
                return false
            }
			if m[sli[len(sli)-1]] == string(v) {
				sli = sli[ : len(sli)-1]
			} else {
				return false
			}
		}
	}
	return len(sli) == 0
}

package main

import "fmt"

func RomanToInt(roman string) int {

	m := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	ks := []string{"CM", "CD", "XC", "XL", "IX", "IV", "M", "D", "C", "L", "X", "V", "I"}

	var num int

	for ;len(roman) > 0;  {
		for _, v := range ks {
			if len(v) == 2 && len(roman)>=2 && v == roman[:2]{
				roman = roman[2:]
				num += m[v]
				break
			} else if len(v) == 1 && v==roman[:1] {
				roman = roman[1:]
				num += m[v]
				break
			}
		}
	}

	return num
}

func main() {
	fmt.Println(RomanToInt("III"))
	fmt.Println(RomanToInt("LVIII"))
}

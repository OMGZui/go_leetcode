package main

import "fmt"

// https://leetcode-cn.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]
		sign := 1
		if temp < last {
			//小数在大数的左边，要减去小数
			sign = -1
		}
		res += sign * temp
		last = temp
	}
	return res
}

func romanToIntRune(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	cur := 0
	pre := 0
	for _, v := range s {
		cur = m[v]
		if cur <= pre {
			sum += cur
		} else {
			sum += cur - 2*pre // 减2倍是因为小的已经加了一次了
		}
		pre = cur
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToIntRune("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToIntRune("MCMXCIV"))
	fmt.Println(romanToIntRune("IV"))
}

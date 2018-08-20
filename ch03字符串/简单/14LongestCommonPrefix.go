/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 09:48
 */
package main

import "fmt"

// n*n*n
func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	// 以最小的字符串为基准遍历
	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			// 如果遇到不相等的直接退出
			if strs[j][i] != byte(r) {
				return strs[j][:i]
				//return short[:i]
			}
		}
	}

	return short
}

// 取最小字符串
func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}
// (n-1)*n
func longestCommonPrefixGood(strs []string) string {
	length := len(strs)
	if length < 1 {
		return ""
	} else if length == 1 {
		return strs[0]
	}
	pos := 0
	// 以第一个为基准
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < length; j++ {
			if len(strs[j]) > i && strs[0][i] == strs[j][i] {
				// 保证最后一个字符串都相等，偏移量加1
				if j == length-1 {
					pos++
				}
			} else {
				return strs[0][:pos]
			}
		}
	}
	return strs[0][:pos]
}

func main() {
	strs := []string{"flight", "flower", "flow", "flo"}
	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefixGood(strs))
}

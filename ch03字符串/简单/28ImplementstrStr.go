/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 17:50
 */
package 简单

import "fmt"

func strStr(haystack string, needle string) int {
	haystackLen, needleLen := len(haystack), len(needle)
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i:(i+needleLen)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaaa", "abb"))
	fmt.Println(strStr("", "abb"))
	fmt.Println(strStr("aaaaaa", ""))
	fmt.Println(strStr("", ""))
}

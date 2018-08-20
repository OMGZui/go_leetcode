/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/14
 * Time: 14:22
 */
package 简单

import "fmt"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if n != 0 {
				return n
			}
			continue
		}
		n++
	}
	return n
}

func main() {
	s := "Hello World"
	s1 := "     "
	s2 := " "
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWord(s1))
	fmt.Println(lengthOfLastWord(s2))
}

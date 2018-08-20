/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/14
 * Time: 15:02
 */
package 简单

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{1}
	}

	// 末尾加一
	digits[length-1]++

	// 处理进位
	for i := length - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	// 处理首位的进位
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	d := []int{1, 2, 3, 4}
	d2 := []int{1, 2, 3, 9}
	d3 := []int{9}
	fmt.Println(plusOne(d))
	fmt.Println(plusOne(d2))
	fmt.Println(plusOne(d3))
}

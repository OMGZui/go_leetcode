/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/23
 * Time: 11:02
 */
package main

import "fmt"

func singleNumber(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	n := 0
	res := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[i] != nums[j] {
				n++
			}
		}
		if n == length-1 {
			res = nums[i]
		}
		n = 0
	}
	return res
}

func singleNumberGood(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	res := 0
	for _, n := range nums {
		// n^n == 0
		// a^b^a^b^a == a
		res ^= n
	}
	return res
}

func main() {
	arr := []int{4, 1, 2, 1, 2, 3, 4, 3, 5}
	fmt.Println(singleNumber(arr))
	fmt.Println(singleNumberGood(arr))
}

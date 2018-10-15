/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/15
 * Time: 11:39
 */
package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	recur(nums, []int{}, &res)
	return res
}

func recur(nums, temp []int, res *[][]int) {
	l := len(nums)
	if l == 0 {
		sort.Ints(temp)
		*res = append(*res, temp)
		return
	}
	recur(nums[:l-1], temp, res)
	recur(nums[:l-1], append([]int{nums[l-1]}, temp...), res)
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subsets(arr))
}

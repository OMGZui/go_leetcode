package main

import "fmt"

// 双重循环
func twoSum(nums []int, target int) []int {
	arr := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr[0] = i
				arr[1] = j
			}
		}
	}
	return arr
}

// 利用Map
func twoSumMap(nums []int, target int) []int {
	m := make(map[int]int)
	rs := make([]int, 2)
	for k, v := range nums {
		w := target - v
		if j, ok := m[w]; ok {
			rs[0] = j
			rs[1] = k
			return rs
		} else {
			m[v] = k
		}
	}
	return rs
}

func main() {
	nums := []int{10, 7, 11, 15, 22}
	target := 33
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumMap(nums, target))
}

/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/24
 * Time: 11:47
 */
package main

import (
	"sort"
	"fmt"
)

func fourSum(nums []int, target int) [][]int {
	length := len(nums)

	if length == 0 {
		return make([][]int, 0)
	}

	// 排序
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-3; i++ {
		// 避免添加重复的结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < target:
					// 已经排好序的情况下，小于0则左侧右移
					l ++
				case s > target:
					// 已经排好序的情况下，大于0则右侧左移
					r --
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					// 为避免重复添加，l 和 r 都需要移动到不同的元素上
					l, r = next(nums, l, r)
				}
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func main() {
	arr := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(arr, 0))
}

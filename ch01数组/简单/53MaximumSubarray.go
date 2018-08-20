/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 20:07
 */
package main

import "fmt"

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	temp := nums[0]
	max := temp
	for i := 1; i < length; i++ {
		// 小于0下标重新开始
		if temp < 0 {
			temp = nums[i]
		} else {
			temp += nums[i]
		}

		// 最大值赋值给max
		if max < temp {
			max = temp
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

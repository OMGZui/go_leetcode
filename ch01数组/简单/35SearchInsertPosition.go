/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 19:38
 */
package 简单

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}
	res := 0
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if nums[i] == target {
			res = i
		} else {
			res ++
		}
	}
	return res
}

// 二分查找
func searchInsertGood(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			return mid
		}
	}
	return low
}

func main() {
	nums := []int{1, 2, 5, 6}
	fmt.Println(searchInsert(nums, 7))
	fmt.Println(searchInsert(nums, 0))
	fmt.Println(searchInsert(nums, 3))
	fmt.Println(searchInsert(nums, 5))

	fmt.Println(searchInsertGood(nums, 5))
	fmt.Println(searchInsertGood(nums, 3))
}

/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 17:17
 */
package 简单

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func removeElementGood(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums3 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums2 := []int{3, 2, 2, 3, 4}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(removeElement(nums2, 3))
	fmt.Println(removeElementGood(nums3, 2))
}

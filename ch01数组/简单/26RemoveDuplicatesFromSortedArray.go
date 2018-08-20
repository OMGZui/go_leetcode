/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 16:28
 */
package 简单

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func removeDuplicatesP(nums []int) (int, []int, []int) {
	if len(nums) == 0 {
		return 0, nums, nums
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1, nums[:i+1], nums
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))

	numsP := []int{0,0,1,1,1,2,2,3,3,4}
	n, num, _num := removeDuplicatesP(numsP)
	fmt.Println(n, num, _num)
}

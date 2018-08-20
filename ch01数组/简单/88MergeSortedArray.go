/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/16
 * Time: 11:29
 */
package 简单

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// 深度复制 nums1
	temp := make([]int, m)
	copy(temp, nums1)

	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		// nums2用完了
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		// temp 用完了
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		// 比较后，放入
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
	return nums1
}

func mergeGood(nums1 []int, m int, nums2 []int, n int) []int {
	for m > 0 && n > 0 {
		a, b := nums1[m-1], nums2[n-1]
		if a > b {
			nums1[m+n-1] = a
			m--
		} else {
			nums1[m+n-1] = b
			n--
		}
	}
	if n > 0 {
		copy(nums1[:n], nums2[:n])
	}
	return nums1
}

func main() {
	//nums1 = [1,2,3,0,0,0], m = 3
	//nums2 = [2,5,6],       n = 3
	m, n := 3, 3
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	fmt.Println(merge(num1, m, num2, n))
	fmt.Println(mergeGood(num1, m, num2, n))
}

/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/31
 * Time: 10:22
 */
package main

import "fmt"

// 选择排序 O(n^2)
func selectSort(arr []int) []int {
	l := len(arr)

	for i := 0; i < l-1; i++ {
		min := i // 假设的最小值
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j // 比最小值都小，那么自己就是最小值
			}
		}
		if i != min { // 选出每一轮的最小值放入arr[i]，得出有序序列
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}

func main() {
	arr := []int{5, 3, 8, 6, 4}
	fmt.Println(selectSort(arr))
}

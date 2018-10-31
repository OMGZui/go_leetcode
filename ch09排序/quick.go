/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/31
 * Time: 11:34
 */
package main

import "fmt"

// 快速排序 O(nlogn)
func quickSort(arr []int, left, right int) {
	if left > right {
		return
	}

	pivot := arr[left] // 基准数
	i := left
	j := right

	for i != j {
		// 右哨兵向左找 大于基准值
		for i < j && arr[j] >= pivot {
			j--
		}
		// 左哨兵向右找 小于基准值
		for i < j && arr[i] <= pivot {
			i++
		}
		// 交换位置
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 相遇则交换基准值
	arr[left] = arr[i]
	arr[i] = pivot

	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}

func sort(arr []int) {
	l := len(arr)
	quickSort(arr, 0, l-1)
}

func main() {
	arr := []int{5, 3, 8, 6, 4}
	sort(arr)
	fmt.Println(arr)
}

/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/31
 * Time: 09:57
 */
package main

import "fmt"

// 冒泡排序 O(n^2)
func bubbleSort(arr []int) []int {
	l := len(arr)

	// 外层控制剩余个数
	for i := l - 1; i > 0; i-- {
		// 内层邻居比较，把大的冒泡
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{5, 3, 8, 6, 4}
	fmt.Println(bubbleSort(arr))
}

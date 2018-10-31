/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/31
 * Time: 11:08
 */
package main

import "fmt"

// 插入排序 O(n^2)
func insertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ { //无序轮次
		temp := arr[i] // 临时值
		for j := i - 1; j >= 0; j-- { // 有序轮次
			if temp < arr[j] { // 将临时值放入有序位置
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{5, 3, 8, 6, 4}
	fmt.Println(insertSort(arr))
}

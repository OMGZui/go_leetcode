/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/9/4
 * Time: 10:42
 */
package main

import "fmt"

func bubble(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
}

func show(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	arr := []int{10, -20, 5, 99, -49, 49}
	bubble(arr)
	show(arr)
}

/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/23
 * Time: 09:41
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

// 两层循环 时间O(n^2) 空间O(1)
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	max := 0
	for x := 0; x < length; x++ {
		for y := x + 1; y < length; y++ {
			temp := prices[y] - prices[x]
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

// 局部最优和全局最优解法 时间O(n) 空间O(1)
func maxProfitGood(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	max := 0
	temp := 0
	for i := 1; i < length; i++ {
		temp = leetcode_tool.Max(temp+prices[i]-prices[i-1], 0)
		max = leetcode_tool.Max(temp, max)
	}

	return max
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	arr2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(arr))
	fmt.Println(maxProfit(arr2))
	fmt.Println(maxProfitGood(arr))
	fmt.Println(maxProfitGood(arr2))
}

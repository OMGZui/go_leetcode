/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/23
 * Time: 09:41
 */
package main

import (
	"fmt"
)

// 贪心算法，总是做出在当前看来是最好的选择，不从整体最优上加以考虑，也就是说，只关心当前最优解
// 局部最优和全局最优解法 时间O(n) 空间O(1)
func maxProfit2(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	max := 0
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	arr2 := []int{7, 6, 4, 3, 1}
	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit2(arr))
	fmt.Println(maxProfit2(arr2))
	fmt.Println(maxProfit2(arr3))
}

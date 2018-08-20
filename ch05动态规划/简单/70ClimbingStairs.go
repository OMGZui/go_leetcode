/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/15
 * Time: 16:26
 */
package main

import "fmt"

// 直接递归 时间O(2^n)
func climbStairs1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

// 备忘录算法 利用map 时间O(n) 空间O(1)
func climbStairs2(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	container := make(map[int]int)
	if container[n] != 0 {
		return container[n]
	} else {
		value := climbStairs2(n-1) + climbStairs2(n-2)
		container[n] = value
		return value
	}
}

// 动态规划  时间O(n)
func climbStairs3(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	a, b := 1, 2
	temp := 0
	for i := 3; i <= n; i++ {
		temp = a + b
		a, b = b, temp
	}
	return temp
}

func main() {
	fmt.Println(climbStairs1(10))
	fmt.Println(climbStairs2(10))
	fmt.Println(climbStairs3(10))
}

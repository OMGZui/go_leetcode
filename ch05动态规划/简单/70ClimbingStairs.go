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
	if n == 0 || n == 1 || n == 2{
		return n
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

// 备忘录算法 利用map 时间O(n) 空间O(n)
func climbStairs2(n int, con map[int]int) int {
	if n == 0 || n == 1 || n == 2{
		return n
	}

	if con[n] != 0 {
		fmt.Println(con[n])
		return con[n]
	} else {
		value := climbStairs2(n-1, con) + climbStairs2(n-2, con)
		con[n] = value
		return value
	}
}

// 动态规划  自底向上 时间O(n)
func climbStairs3(n int) int {
	if n == 0 || n == 1 || n == 2{
		return n
	}

	// 代表上次和上上次的结果，实际上就是F(n-1),F(n-2)
	a, b := 1, 2
	// 代表F(n)
	temp := 0
	for i := 3; i <= n; i++ {
		temp = a + b
		a, b = b, temp
	}
	return temp
}

func main() {
	fmt.Println(climbStairs1(10))
	container := make(map[int]int)
	fmt.Println(climbStairs2(10, container))
	fmt.Println(climbStairs3(10))
}

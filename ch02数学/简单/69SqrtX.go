/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/15
 * Time: 15:41
 */
package 简单

import "fmt"

func mySqrt(x int) int {
	res := x

	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
}
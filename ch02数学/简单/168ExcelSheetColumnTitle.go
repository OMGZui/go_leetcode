/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/23
 * Time: 14:55
 */
package main

import "fmt"

func convertToTitle(n int) string {
	res := ""

	for n > 0 {
		n--
		res = string(byte(n%26)+'A') + res
		n /= 26
	}

	return res
}

func main() {
	fmt.Println(convertToTitle(701))
}

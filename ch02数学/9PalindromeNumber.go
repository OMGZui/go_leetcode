package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-11))
}

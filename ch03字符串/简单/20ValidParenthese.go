/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/13
 * Time: 10:30
 */
package 简单

import (
	"fmt"
	"container/list"
)

var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

type stack []rune

func isValid(str string) bool {
	s := new(stack)

	for _, b := range str {
		switch b {
		case '(', '[', '{':
			s.push(b)
		case ')', ']', '}':
			if r, ok := s.pop(); !ok || r != matching[b] {
				// !ok 说明“（[{”的数量，小于")]}"的数量
				return false
			}
		}
	}

	// len(*s) > 0 说明"([{"的数量，大于")]}"的数量
	if len(*s) > 0 {
		return false
	}

	return true
}

// 入栈
func (s *stack) push(b rune) {
	*s = append(*s, b)
}

// 出栈
func (s *stack) pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}

func isValidList(s string) bool {
	m := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	stack := list.New()

	for _, v := range s {
		if stack.Len() == 0 {
			stack.PushFront(string(v))
			continue
		}

		if stack.Front().Value.(string) == m[string(v)] {
			stack.Remove(stack.Front())
			continue
		}

		stack.PushFront(string(v))
	}

	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValidList("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValidList("([)]"))
}

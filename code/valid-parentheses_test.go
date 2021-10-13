package code

import (
	"testing"
)

/*
有效的括号（简单）
http://leetcode-cn.com/problems/valid-parentheses
*/
func TestIsValid(t *testing.T) {
	s := "{[]}"
	t.Log(isValid(s))
}

// 栈
func isValid(s string) bool {
	lens := len(s)
	if lens <= 1 {
		return false
	}

	var (
		sbyt  = []byte(s)
		stack = make([]byte, lens)
		top   = -1
	)
	for i := 0; i < lens; i++ {
		if top < 0 {
			stack[0] = sbyt[i]
			top = 0
			continue
		}
		if (stack[top] == '(' && sbyt[i] == ')') || (stack[top] == '[' && sbyt[i] == ']') || (stack[top] == '{' && sbyt[i] == '}') {
			top--
		} else {
			top++
			stack[top] = sbyt[i]
		}
	}
	return top < 0
}

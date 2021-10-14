package code

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
字符串解码（中等）
http://leetcode-cn.com/problems/decode-string
*/
func TestDecodeString(t *testing.T) {
	var s = "3[a]2[c]"
	t.Log(decodeString(s))
}

/* 双栈
入栈时机：遇到[。意味着要解决内部，外部的数字和字母，去栈里等。
当遇到[，已经扫描的数字就是“倍数”，入栈暂存
当遇到[，已经扫描的字母也入栈等待，括号里的解码完了，一起参与构建字符串。

出栈时机：遇到]。内层的扫描完了，栈顶元素可以出栈了，共同参与子串的构建。
栈顶就是最近遇到的“倍数”和字母
*/
func decodeString(s string) string {
	var (
		lens     = len(s) / 2
		numStack = make([]int, lens)
		numTop   = -1
		strStack = make([]string, lens)
		strTop   = -1
		num      = 0
		str      = ""
	)

	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			numTop++
			numStack[numTop] = num
			num = 0

			strTop++
			strStack[strTop] = str
			str = ""
		} else if char == ']' {
			count := numStack[numTop]
			numTop--
			st := strStack[strTop]
			strTop--
			str = st + strings.Repeat(str, count)
		} else {
			str += string(char)
		}
		fmt.Println(str)
	}

	return str
}

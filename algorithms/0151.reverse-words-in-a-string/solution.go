package main

import (
	"fmt"
	"reflect"
)

func reverseWords(s string) string {
	b := []byte(s)

	// 移除前面、中间、后面存在的多余空格
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' { // 复制逻辑
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[0:slow]

	// 翻转整个字符串
	reverse(b)
	// 翻转每个单词
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b[last:i])
			last = i + 1
		}
	}
	return string(b)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

type Example struct {
	s   string
	ans string
}

func main() {
	examples := []Example{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}
	for i, e := range examples {
		ans := reverseWords(e.s)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}

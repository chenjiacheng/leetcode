package main

import (
	"fmt"
	"reflect"
)

func longestCommonPrefix(strs []string) string {
	ans := ""
	str := strs[0]
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) == i || str[i] != strs[j][i] {
				return ans
			}
		}
		ans += string(str[i])
	}
	return ans
}

type Example struct {
	nums []string
	ans  string
}

func main() {
	examples := []Example{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "a"}, "a"},
	}
	for i, e := range examples {
		ans := longestCommonPrefix(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}

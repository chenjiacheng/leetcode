package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	path []string
	res  []string
)

func restoreIpAddresses(s string) []string {
	path, res = make([]string, 0, len(s)), make([]string, 0)
	dfs(s, 0)
	return res
}
func dfs(s string, start int) {
	if len(path) == 4 { // 够四段后就不再继续往下递归
		if start == len(s) {
			str := strings.Join(path, ".")
			res = append(res, str)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' { // 含有前导 0，无效
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path = append(path, str) // 符合条件的就进入下一层
			dfs(s, i+1)
			path = path[:len(path)-1]
		} else { // 如果不满足条件，再往后也不可能满足条件，直接退出
			break
		}
	}
}

type Example struct {
	s   string
	ans []string
}

func main() {
	examples := []Example{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"0000", []string{"0.0.0.0"}},
		{"101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}
	for i, e := range examples {
		ans := restoreIpAddresses(e.s)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}

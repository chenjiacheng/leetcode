package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	mp := map[string][]string{}

	for _, str := range strs {
		sorted := sortString(str)
		mp[sorted] = append(mp[sorted], str)
	}

	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}

	return res
}

func sortString(str string) string {
	// 将字符串转换为字符切片
	chars := strings.Split(str, "")

	// 使用 sort 包对字符切片进行排序
	sort.Strings(chars)

	// 将排序后的字符切片重新组合为字符串
	return strings.Join(chars, "")
}

type Example struct {
	strs []string
	ans  [][]string
}

func main() {
	examples := []Example{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for i, e := range examples {
		ans := groupAnagrams(e.strs)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}

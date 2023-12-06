package leetcode0049

import (
	"sort"
	"strings"
)

// 解题思路：
// 哈希表

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

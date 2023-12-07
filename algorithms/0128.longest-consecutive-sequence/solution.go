package leetcode0128

// 解题思路：
// 哈希表

func longestConsecutive(nums []int) int {
	// 把值存储到hash表
	mp := map[int]bool{}
	for _, num := range nums {
		mp[num] = true
	}

	ans := 0 // 最长长度
	for num := range mp {
		// 如果是起点就计算长度
		if _, ok := mp[num-1]; !ok {
			x := num + 1
			for mp[x] {
				x++
			}
			if x-num > ans {
				ans = x - num
			}
		}
	}

	return ans
}

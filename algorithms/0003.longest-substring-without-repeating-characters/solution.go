package leetcode0003

// 解题思路：
// 滑动窗口

func lengthOfLongestSubstring(s string) int {
	// 初始化 左指针，右指针，当前长度，最大长度
	left, right, length, maxLength := 0, 0, 0, 0

	// 记录每个字符是否出现过
	m := map[byte]int{}

	n := len(s)

	for right < n {
		if _, ok := m[s[right]]; !ok {
			// 如果字符没有出现过，记录字符并移动右指针
			m[s[right]] = 0
			right++
			length++
			if length > maxLength {
				maxLength++
			}
		} else {
			// 如果字符出现过，不断移动左指针
			for {
				if _, ok := m[s[right]]; !ok {
					break
				}

				delete(m, s[left])
				left++
				length--
			}

			// 重新记录字符并移动右指针
			m[s[right]] = 0
			right++
			length++
		}
	}

	return maxLength
}

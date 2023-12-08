package leetcode0011

// 解题思路：
// 双指针

func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0

	for left < right {

		res := 0

		// 当右边大于左边，需要移动左边，反之移动右边
		if height[right] > height[left] {
			res = height[left] * (right - left)
			left++
		} else {
			res = height[right] * (right - left)
			right--
		}

		if res > ans {
			ans = res
		}
	}

	return ans
}

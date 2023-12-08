package leetcode0283

// 解题思路：
// 双指针

func moveZeroes(nums []int) {
	left := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}

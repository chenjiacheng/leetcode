package leetcode0001

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, num := range nums {
		res, ok := myMap[target-num]
		if ok {
			return []int{res, i}
		} else {
			myMap[num] = i
		}
	}
	return []int{}
}

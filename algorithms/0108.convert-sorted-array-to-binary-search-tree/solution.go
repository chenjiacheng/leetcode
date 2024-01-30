package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	k := l / 2
	return &TreeNode{Val: nums[k], Left: sortedArrayToBST(nums[:k]), Right: sortedArrayToBST(nums[k+1:])}
}

type Example struct {
	nums []int
	ans  *TreeNode
}

func main() {
	examples := []Example{
		{[]int{-10, -3, 0, 5, 9}, &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}}},
		{[]int{1, 3}, &TreeNode{Val: 3, Left: &TreeNode{Val: 1}}},
	}
	for i, e := range examples {
		ans := sortedArrayToBST(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}

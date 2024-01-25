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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	k, v := 0, 0
	for i, num := range nums {
		if num > v {
			k, v = i, num
		}
	}
	return &TreeNode{
		Val:   v,
		Left:  constructMaximumBinaryTree(nums[0:k]),
		Right: constructMaximumBinaryTree(nums[k+1:]),
	}
}

type Example struct {
	nums []int
	ans  *TreeNode
}

func main() {
	examples := []Example{
		{[]int{3, 2, 1, 6, 0, 5}, &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 0}}}},
		{[]int{3, 2, 1}, &TreeNode{Val: 3, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}},
	}
	for i, e := range examples {
		ans := constructMaximumBinaryTree(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}

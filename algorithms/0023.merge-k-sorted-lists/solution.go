package leetcode0023

// 解题思路：
// 分治合并

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	left, right := 0, length-1
	for right > 0 {
		for left < right {
			// 头尾合并，并移动左和右指针
			lists[left] = mergeTwoLists(lists[left], lists[right])
			left++
			right--
		}
		left = 0
	}

	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var node *ListNode
	if l1.Val >= l2.Val {
		node = l2
		node.Next = mergeTwoLists(l1, l2.Next)
	} else {
		node = l1
		node.Next = mergeTwoLists(l1.Next, l2)
	}

	return node
}

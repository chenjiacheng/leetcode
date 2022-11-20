package leetcode0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)

	// 删除头节点
	if length-n == 0 {
		head = head.Next
		return head
	}

	dummy := head
	i := 1
	for dummy != nil {
		if length-i == n {
			dummy.Next = dummy.Next.Next
			break
		}
		dummy = dummy.Next
		i++
	}

	return head
}

func getLength(l *ListNode) int {
	length := 1
	for l.Next != nil {
		length++
		l = l.Next
	}

	return length
}

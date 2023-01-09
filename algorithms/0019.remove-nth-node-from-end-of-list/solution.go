package leetcode0019

// 解题思路：
// 1. 计算链表长度
// 2. 遍历查找对应位置的前一个
// 3. 注意另外处理删除头节点

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

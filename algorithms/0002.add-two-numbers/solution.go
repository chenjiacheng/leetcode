package leetcode0002

// 解题思路：
// 不对齐补零，若链表不为空则用 sum(代表每个位的和的结果)加上，考虑进位。

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r, p *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		sum, carry = sum%10, sum/10
		if r == nil {
			r = &ListNode{Val: sum}
			p = r
		} else {
			p.Next = &ListNode{Val: sum}
			p = p.Next
		}
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return r
}

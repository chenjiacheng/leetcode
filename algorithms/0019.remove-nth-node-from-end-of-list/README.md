## 19. 删除链表的倒数第 N 个结点

`中等`

## 题目描述

给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

**示例 1：**

![img](README.assets/remove_ex1.jpg)

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]
```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]
```

**提示：**

- 链表中结点的数目为 `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**进阶：**你能尝试使用一趟扫描实现吗？

## 解题思路

- 计算链表长度
- 遍历查找对应位置的前一个
- 注意另外处理删除头节点

## 代码实现

```go
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
```


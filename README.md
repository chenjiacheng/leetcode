# 力扣（LeetCode）

## 调试

进入题解所在目录执行以下命令：

```bash
# Go
go run solution.go
```

每通过一个用例会输出：`PASS: CASE XXX`，未通过会输出：`FAIL: CASE XXX`。

## 题解

### 算法

| 题目 | 相关标签 | 题解 | 难度 |
| ---- | ---- | ---- | ---- |
| [1. 两数之和](https://leetcode.cn/problems/two-sum/) | `数组` `哈希表` | [Go](./algorithms/0001.two-sum/solution.go) | 简单 |
| [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/) | `递归` `链表` `数学` | [Go](./algorithms/0002.add-two-numbers/solution.go) | 中等 |
| [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) | `哈希表` `字符串` `滑动窗口` | [Go](./algorithms/0003.longest-substring-without-repeating-characters/solution.go) | 中等 |
| [11. 盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/) | `贪心` `数组` `双指针` | [Go](./algorithms/0011.container-with-most-water/solution.go) | 中等 |
| [15. 三数之和](https://leetcode.cn/problems/3sum/) | `数组` `双指针` `排序` | [Go](./algorithms/0015.3sum/solution.go) | 中等 |
| [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) | `链表` `双指针` | [Go](./algorithms/0019.remove-nth-node-from-end-of-list/solution.go) | 中等 |
| [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/) | `栈` `字符串` | [Go](./algorithms/0020.valid-parentheses/solution.go) | 简单 |
| [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/) | `递归` `链表` | [Go](./algorithms/0021.merge-two-sorted-lists/solution.go) | 简单 |
| [23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/) | `链表` `分治` `堆（优先队列）` `归并排序` | [Go](./algorithms/0023.merge-k-sorted-lists/solution.go) | 困难 |
| [24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/) | `递归` `链表` | [Go](./algorithms/0024.swap-nodes-in-pairs/solution.go) | 中等 |
| [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/) | `数组` `双指针` | [Go](./algorithms/0026.remove-duplicates-from-sorted-array/solution.go) | 简单 |
| [27. 移除元素](https://leetcode.cn/problems/remove-element/) | `数组` `双指针` | [Go](./algorithms/0027.remove-element/solution.go) | 简单 |
| [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/) | `数组` `哈希表` `字符串` `排序` | [Go](./algorithms/0049.group-anagrams/solution.go) | 中等 |
| [54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/) | `数组` `矩阵` `模拟` | [Go](./algorithms/0054.spiral-matrix/solution.go) | 中等 |
| [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/) | `数组` `矩阵` `模拟` | [Go](./algorithms/0059.spiral-matrix-ii/solution.go) | 中等 |
| [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/) | `并查集` `数组` `哈希表` | [Go](./algorithms/0128.longest-consecutive-sequence/solution.go) | 中等 |
| [202. 快乐数](https://leetcode.cn/problems/happy-number/) | `哈希表` `数学` `双指针` | [Go](./algorithms/0202.happy-number/solution.go) | 简单 |
| [203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/) | `递归` `链表` | [Go](./algorithms/0203.remove-linked-list-elements/solution.go) | 简单 |
| [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/) | `递归` `链表` | [Go](./algorithms/0206.reverse-linked-list/solution.go) | 简单 |
| [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) | `数组` `二分查找` `前缀和` `滑动窗口` | [Go](./algorithms/0209.minimum-size-subarray-sum/solution.go) | 中等 |
| [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/) | `哈希表` `字符串` `排序` | [Go](./algorithms/0242.valid-anagram/solution.go) | 简单 |
| [283. 移动零](https://leetcode.cn/problems/move-zeroes/) | `数组` `双指针` | [Go](./algorithms/0283.move-zeroes/solution.go) | 简单 |
| [349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/) | `数组` `哈希表` `双指针` `二分查找` `排序` | [Go](./algorithms/0349.intersection-of-two-arrays/solution.go) | 简单 |
| [704. 二分查找](https://leetcode.cn/problems/binary-search/) | `数组` `二分查找` | [Go](./algorithms/0704.binary-search/solution.go) | 简单 |
| [844. 比较含退格的字符串](https://leetcode.cn/problems/backspace-string-compare/) | `栈` `双指针` `字符串` `模拟` | [Go](./algorithms/0844.backspace-string-compare/solution.go) | 简单 |
| [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/) | `数组` `双指针` `排序` | [Go](./algorithms/0977.squares-of-a-sorted-array/solution.go) | 简单 |

### 数据库

| 题目 | 题解 | 难度 |
| ---- | ---- | ---- |
| [175. 组合两个表](https://leetcode.cn/problems/combine-two-tables/) | [MySQL](./database/0175.combine-two-tables/solution.sql) | 简单 |
| [178. 分数排名](https://leetcode.cn/problems/rank-scores/) | [MySQL](./database/0178.rank-scores/solution.sql) | 中等 |


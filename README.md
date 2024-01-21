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
| [9. 回文数](https://leetcode.cn/problems/palindrome-number/) | `数学` | [Go](./algorithms/0009.palindrome-number/solution.go) | 简单 |
| [11. 盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/) | `贪心` `数组` `双指针` | [Go](./algorithms/0011.container-with-most-water/solution.go) | 中等 |
| [14. 最长公共前缀](https://leetcode.cn/problems/longest-common-prefix/) | `字典树` `字符串` | [Go](./algorithms/0014.longest-common-prefix/solution.go) | 简单 |
| [15. 三数之和](https://leetcode.cn/problems/3sum/) | `数组` `双指针` `排序` | [Go](./algorithms/0015.3sum/solution.go) | 中等 |
| [18. 四数之和](https://leetcode.cn/problems/4sum/) | `数组` `双指针` `排序` | [Go](./algorithms/0018.4sum/solution.go) | 中等 |
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
| [94. 二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/) | `栈` `树` `深度优先搜索` `二叉树` | [Go](./algorithms/0094.binary-tree-inorder-traversal/solution.go) | 简单 |
| [101. 对称二叉树](https://leetcode.cn/problems/symmetric-tree/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0101.symmetric-tree/solution.go) | 简单 |
| [102. 二叉树的层序遍历](https://leetcode.cn/problems/binary-tree-level-order-traversal/) | `树` `广度优先搜索` `二叉树` | [Go](./algorithms/0102.binary-tree-level-order-traversal/solution.go) | 中等 |
| [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0104.maximum-depth-of-binary-tree/solution.go) | 简单 |
| [107. 二叉树的层序遍历 II](https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/) | `树` `广度优先搜索` `二叉树` | [Go](./algorithms/0107.binary-tree-level-order-traversal-ii/solution.go) | 中等 |
| [110. 平衡二叉树](https://leetcode.cn/problems/balanced-binary-tree/) | `树` `深度优先搜索` `二叉树` | [Go](./algorithms/0110.balanced-binary-tree/solution.go) | 简单 |
| [111. 二叉树的最小深度](https://leetcode.cn/problems/minimum-depth-of-binary-tree/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0111.minimum-depth-of-binary-tree/solution.go) | 简单 |
| [116. 填充每个节点的下一个右侧节点指针](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/) | `树` `深度优先搜索` `广度优先搜索` `链表` `二叉树` | [Go](./algorithms/0116.populating-next-right-pointers-in-each-node/solution.go) | 中等 |
| [117. 填充每个节点的下一个右侧节点指针 II](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/) | `树` `深度优先搜索` `广度优先搜索` `链表` `二叉树` | [Go](./algorithms/0117.populating-next-right-pointers-in-each-node-ii/solution.go) | 中等 |
| [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/) | `并查集` `数组` `哈希表` | [Go](./algorithms/0128.longest-consecutive-sequence/solution.go) | 中等 |
| [144. 二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/) | `栈` `树` `深度优先搜索` `二叉树` | [Go](./algorithms/0144.binary-tree-preorder-traversal/solution.go) | 简单 |
| [145. 二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/) | `栈` `树` `深度优先搜索` `二叉树` | [Go](./algorithms/0145.binary-tree-postorder-traversal/solution.go) | 简单 |
| [150. 逆波兰表达式求值](https://leetcode.cn/problems/evaluate-reverse-polish-notation/) | `栈` `数组` `数学` | [Go](./algorithms/0150.evaluate-reverse-polish-notation/solution.go) | 中等 |
| [151. 反转字符串中的单词](https://leetcode.cn/problems/reverse-words-in-a-string/) | `双指针` `字符串` | [Go](./algorithms/0151.reverse-words-in-a-string/solution.go) | 中等 |
| [199. 二叉树的右视图](https://leetcode.cn/problems/binary-tree-right-side-view/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0199.binary-tree-right-side-view/solution.go) | 中等 |
| [202. 快乐数](https://leetcode.cn/problems/happy-number/) | `哈希表` `数学` `双指针` | [Go](./algorithms/0202.happy-number/solution.go) | 简单 |
| [203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/) | `递归` `链表` | [Go](./algorithms/0203.remove-linked-list-elements/solution.go) | 简单 |
| [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/) | `递归` `链表` | [Go](./algorithms/0206.reverse-linked-list/solution.go) | 简单 |
| [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) | `数组` `二分查找` `前缀和` `滑动窗口` | [Go](./algorithms/0209.minimum-size-subarray-sum/solution.go) | 中等 |
| [222. 完全二叉树的节点个数](https://leetcode.cn/problems/count-complete-tree-nodes/) | `位运算` `树` `二分查找` `二叉树` | [Go](./algorithms/0222.count-complete-tree-nodes/solution.go) | 简单 |
| [226. 翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0226.invert-binary-tree/solution.go) | 简单 |
| [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) | `队列` `数组` `滑动窗口` `单调队列` `堆（优先队列）` | [Go](./algorithms/0239.sliding-window-maximum/solution.go) | 困难 |
| [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/) | `哈希表` `字符串` `排序` | [Go](./algorithms/0242.valid-anagram/solution.go) | 简单 |
| [283. 移动零](https://leetcode.cn/problems/move-zeroes/) | `数组` `双指针` | [Go](./algorithms/0283.move-zeroes/solution.go) | 简单 |
| [344. 反转字符串](https://leetcode.cn/problems/reverse-string/) | `双指针` `字符串` | [Go](./algorithms/0344.reverse-string/solution.go) | 简单 |
| [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/) | `数组` `哈希表` `分治` `桶排序` `计数` `快速选择` `排序` `堆（优先队列）` | [Go](./algorithms/0347.top-k-frequent-elements/solution.go) | 中等 |
| [349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/) | `数组` `哈希表` `双指针` `二分查找` `排序` | [Go](./algorithms/0349.intersection-of-two-arrays/solution.go) | 简单 |
| [383. 赎金信](https://leetcode.cn/problems/ransom-note/) | `哈希表` `字符串` `计数` | [Go](./algorithms/0383.ransom-note/solution.go) | 简单 |
| [429. N 叉树的层序遍历](https://leetcode.cn/problems/n-ary-tree-level-order-traversal/) | `树` `广度优先搜索` | [Go](./algorithms/0429.n-ary-tree-level-order-traversal/solution.go) | 中等 |
| [454. 四数相加 II](https://leetcode.cn/problems/4sum-ii/) | `数组` `哈希表` | [Go](./algorithms/0454.4sum-ii/solution.go) | 中等 |
| [515. 在每个树行中找最大值](https://leetcode.cn/problems/find-largest-value-in-each-tree-row/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0515.find-largest-value-in-each-tree-row/solution.go) | 中等 |
| [541. 反转字符串 II](https://leetcode.cn/problems/reverse-string-ii/) | `双指针` `字符串` | [Go](./algorithms/0541.reverse-string-ii/solution.go) | 简单 |
| [637. 二叉树的层平均值](https://leetcode.cn/problems/average-of-levels-in-binary-tree/) | `树` `深度优先搜索` `广度优先搜索` `二叉树` | [Go](./algorithms/0637.average-of-levels-in-binary-tree/solution.go) | 简单 |
| [704. 二分查找](https://leetcode.cn/problems/binary-search/) | `数组` `二分查找` | [Go](./algorithms/0704.binary-search/solution.go) | 简单 |
| [844. 比较含退格的字符串](https://leetcode.cn/problems/backspace-string-compare/) | `栈` `双指针` `字符串` `模拟` | [Go](./algorithms/0844.backspace-string-compare/solution.go) | 简单 |
| [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/) | `数组` `双指针` `排序` | [Go](./algorithms/0977.squares-of-a-sorted-array/solution.go) | 简单 |
| [1047. 删除字符串中的所有相邻重复项](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/) | `栈` `字符串` | [Go](./algorithms/1047.remove-all-adjacent-duplicates-in-string/solution.go) | 简单 |

### 数据库

| 题目                                                         | 题解                                                        | 难度 |
| ------------------------------------------------------------ | ----------------------------------------------------------- | ---- |
| [175. 组合两个表](https://leetcode.cn/problems/combine-two-tables/) | [MySQL](./database/0175.combine-two-tables/solution.sql)    | 简单 |
| [176. 第二高的薪水](https://leetcode.cn/problems/second-highest-salary/) | [MySQL](./database/0176.second-highest-salary/solution.sql) | 中等 |
| [177. 第N高的薪水](https://leetcode.cn/problems/nth-highest-salary/) | [MySQL](./database/0177.nth-highest-salary/solution.sql)    | 中等 |
| [178. 分数排名](https://leetcode.cn/problems/rank-scores/)   | [MySQL](./database/0178.rank-scores/solution.sql)           | 中等 |


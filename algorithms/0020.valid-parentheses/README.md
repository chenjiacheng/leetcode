# 20. 有效的括号

`简单`

## 题目描述

给定一个只包括 `'('`，`')'`，`'{'`，`'}'`，`'['`，`']'` 的字符串 `s` ，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

**示例 1：**

```
输入：s = "()"
输出：true
```

**示例 2：**

```
输入：s = "()[]{}"
输出：true
```

**示例 3：**

```
输入：s = "(]"
输出：false
```

**提示：**

- `1 <= s.length <= 104`
- `s` 仅由括号 `'()[]{}'` 组成

## 解题思路

- 栈

## 代码实现

```go
func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []byte

	for _, v := range s {
		switch v {
		// 如果是左边就入栈
		case '(', '{', '[':
			stack = append(stack, byte(v))
		// 如果是右边则从栈中取出一个做对比
		case ')', '}', ']':
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != byte(v) {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
```

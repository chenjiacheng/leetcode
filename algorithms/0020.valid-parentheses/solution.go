package leetcode0020

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

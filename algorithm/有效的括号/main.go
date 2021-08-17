package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	hash := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for n := 0; n < len(s); n++ {
		if hash[s[n]] > 0 {
			if len(stack) == 0 || hash[s[n]] != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]	// 如果匹配弹出栈顶
			}
		} else {
			stack = append(stack, s[n])
		}
	}
	return len(stack) == 0 // 如果全部匹配那么栈为空
}

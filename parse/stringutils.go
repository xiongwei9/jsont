package parse

// 从 start 位开始过滤空白字符，返回非遇到的第一个非空白字符索引
// -1 表示没有非空白字符
func passSpace(str []rune, start int) int {
	if start < 0 {
		return -1
	}
	strLen := len(str)
	for i := start; i < strLen; i++ {
		if str[i] != ' ' { // TODO: newline/tab...
			return i
		}
	}
	return -1
}

func getQuoteContent(str []rune, start int) ([]rune, int) {
	if start < 0 || str[start] != '"' {
		return nil, -1
	}
	end := -1
	for i := start + 1; i < len(str); i++ {
		if str[i] == '"' {
			end = i
			break
		}
	}
	var res []rune
	if end > 0 && start+1 < end-1 {
		res = str[start+1 : end]
	}
	return res, end
}

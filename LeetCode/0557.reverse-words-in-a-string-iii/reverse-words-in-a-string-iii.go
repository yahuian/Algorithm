package _0557

func reverseWords(s string) string {
	var words [][]byte
	var word string

	s += " " // 给s后面加一个空格，方便我们拆分字符串
	// 按照空格拆分字符串为单词数组
	for _, v := range s {
		if string(v) != " " {
			word += string(v)
		} else {
			words = append(words, []byte(word))
			word = ""
		}
	}

	// 反转单个单词
	for _, v := range words {
		for i := 0; i < len(v)/2; i++ {
			v[i], v[len(v)-i-1] = v[len(v)-i-1], v[i]
		}
	}

	var result string
	// 合并为一个字符串
	for _, v := range words {
		result += string(v) + " "
	}

	return result[:len(result)-1] // 去除末尾的空格
}

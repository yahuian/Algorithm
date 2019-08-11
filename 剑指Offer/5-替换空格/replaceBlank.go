package main

func main() {
	str1 := "we are happy."
}

func func1(aStr string) string {
	// 统计aStr中空格的数量
	count := 0
	for i := 0; i < len(aStr); i++ {
		if aStr[i] == ' ' {
			count++
		}
	}
	temp := make([]char, len(aStr)+2*count) // 一个空格被%20三个字符所替代

}

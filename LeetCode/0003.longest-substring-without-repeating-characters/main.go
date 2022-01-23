package main

func main() {
}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// 暴力法，两层循环，O(n^2)
func lengthOfLongestSubstringOne(s string) int {
	var max int
	for i := range s {
		var tmp int
		m := make(map[string]struct{}, 0)
		for _, v2 := range s[i:] {
			key := string(v2)
			if _, ok := m[key]; ok {
				break
			}
			m[key] = struct{}{}
			tmp++
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

package _0242

func isAnagram(s string, t string) bool {
	map1 := map[string]int{}
	map2 := map[string]int{}

	for _, v1 := range s {
		map1[string(v1)]++
	}

	for _, v2 := range t {
		map2[string(v2)]++
	}

	// 判断map1和map2是否相等
	if len(map1) != len(map2) {
		return false
	}
	for key1, value1 := range map1 {
		value, ok := map2[key1]
		if ok == false {
			return false
		} else if value1 != value {
			return false
		}
	}

	return true
}

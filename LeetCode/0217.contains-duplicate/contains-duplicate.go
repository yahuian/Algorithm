package _0217

func containsDuplicate(nums []int) bool {
	resMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := resMap[nums[i]]; ok == true {
			return true
		} else {
			resMap[nums[i]] = 0
		}
	}
	return false
}

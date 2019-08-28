package _0001

func twoSum(nums []int, target int) []int {
	aMap := make(map[int]int, len(nums))
	for index, value := range nums {
		y := target - value
		if v, ok := aMap[y]; ok == true {
			return []int{v, index}
		}
		aMap[value] = index
	}
	return []int{}
}

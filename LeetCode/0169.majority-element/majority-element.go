package _0169

// 摩尔投票法
// 注意审题，本题中的众数指的是在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 而一般意义上的众数指的是一组数据中出现次数最多的数据值。
func majorityElement(nums []int) int {
	mode, count := nums[0], 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == mode {
			count++
		} else {
			count--
		}
		if count == 0 {
			mode = nums[i]
			count = 1
		}
	}
	return mode
}

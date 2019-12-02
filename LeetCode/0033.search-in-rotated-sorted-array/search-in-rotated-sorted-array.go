package _0033

// 本题思路：
// 先二分法找数组的旋转点
// 再二分法找target
// 注意题干中说到的O(nlgn)暗示我们二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min := min(nums)

	if nums[min] == target {
		return min
	}

	// 数组没有旋转
	if min == 0 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	// 在后一半升序数组中找
	if target < nums[0] {
		return binarySearch(nums, min, len(nums)-1, target)
	} else { // 在前一半升序数组中找
		return binarySearch(nums, 0, min, target)
	}
}

// 寻找旋转点，即数组的最小值
func min(nums []int) int {
	p1 := 0
	p2 := len(nums) - 1
	// 将mid设置为p1,一旦发现第一个数小于最后一个数，说明数组旋转了0位
	// 即没旋转，可以直接返回答案nums[p1],即第一个数
	mid := p1
	for nums[p1] >= nums[p2] {
		// 说明p1指向了第一个递增数组的尾部
		// p2指向了第二个递增数组的头部，即最小值
		if p2-p1 == 1 {
			mid = p2
			break
		}

		mid = (p1 + p2) / 2

		// 如果p1,mid,p2三者指向的数字相同，则没法判断最小值位于什么位置，只能遍历查找
		if nums[p1] == nums[mid] && nums[mid] == nums[p2] {
			res := p1
			for i := p1; i <= p2; i++ {
				if nums[res] > nums[i] {
					res = i
				}
			}
			return res
		}

		if nums[mid] >= nums[p1] { // 说明mid位于第一个递增数组
			p1 = mid
		} else if nums[mid] <= nums[p2] { // 说明mid位于第二个递增数组
			p2 = mid
		}

	}
	return mid
}

// 二分查找
func binarySearch(nums []int, i, j, t int) int {
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

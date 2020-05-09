package binarysearch

// 查找第一个值等于给定值的元素
func BinarySearch2(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] > val {
			hi = mid - 1
		} else if nums[mid] < val {
			lo = mid + 1
		} else {
			if mid == 0 || nums[mid - 1] != val {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个等于给定值的元素
func BinarySearch3(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] > val {
			hi = mid - 1
		} else if nums[mid] < val {
			lo = mid + 1
		} else {
			if mid == len(nums) - 1 || nums[mid + 1] != val {
				return mid
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearch4(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] >= val {
			if mid == 0 || nums[mid - 1] < val {
				return mid
			} else {
				hi = mid - 1
			}
		} else {
			lo = mid + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func BinarySearch5(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] <= val {
			if mid == len(nums) - 1 || nums[mid + 1] > val {
				return mid
			} else {
				lo = mid + 1
			}
		} else {
			lo = mid - 1
		}
	}
	return -1
}
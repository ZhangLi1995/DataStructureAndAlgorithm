package binarysearch

// 循环版本
func BinarySearch(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := lo + (hi - lo) >> 1
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

// 递归版本
func RecursionBinarySearch(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums) - 1
	var recursion func(nums []int, lo, hi, val int) int
	recursion = func(nums []int, lo, hi, val int) int {
		if lo > hi {
			return -1
		}
		mid := lo + (hi - lo) >> 1
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			return recursion(nums, lo, mid-1, val)
		}  else {
			return recursion(nums, lo+1, hi, val)
		}
	}
	return recursion(nums, lo, hi, val)
}

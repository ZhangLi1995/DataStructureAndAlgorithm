package BinarySearch;

/**
 * @author ZhangLi
 * 2020/5/9 10:08
 */
public class BinarySearch2 {

    // 查找第一个值等于给定值的元素
    public static int binarySearch(int[] nums, int val) {
        if (nums == null || nums.length == 0) {
            return -1;
        }
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (nums[mid] > val) {
                hi = mid - 1;
            } else if (nums[mid] < val) {
                lo = mid + 1;
            } else {
                if (mid == 0 || nums[mid - 1] != val) {
                    return mid;
                } else {
                    hi = mid - 1;
                }
            }
        }
        return -1;
    }

    // 查找最后一个等于给定值的元素
    public static int binarySearch2(int[] nums, int val) {
        if (nums == null || nums.length == 0) {
            return -1;
        }
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (nums[mid] > val) {
                hi = mid - 1;
            } else if (nums[mid] < val) {
                lo = mid + 1;
            } else {
                if (mid == nums.length - 1 || nums[mid + 1] != val) {
                    return mid;
                } else {
                    lo = mid + 1;
                }
            }
        }
        return -1;
    }

    // 查找第一个大于等于给定值的元素
    public static int binarySearc3(int[] nums, int val) {
        if (nums == null || nums.length == 0) {
            return -1;
        }
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (nums[mid] >= val) {
                if (mid == 0 || nums[mid - 1] < val) {
                    return mid;
                }
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return -1;
    }

    // 查找最后一个小于等于给定值的元素
    public static int binarySearc4(int[] nums, int val) {
        if (nums == null || nums.length == 0) {
            return -1;
        }
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (nums[mid] > val) {
                hi = mid - 1;
            } else {
                if (mid == nums.length - 1 || nums[mid + 1] > val) {
                    return mid;
                }
                lo = mid + 1;
            }
        }
        return -1;
    }
}

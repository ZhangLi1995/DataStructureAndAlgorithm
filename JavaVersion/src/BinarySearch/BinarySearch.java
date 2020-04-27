package BinarySearch;

/**
 * @author ZhangLi
 * 2020/4/27 10:48
 */
public class BinarySearch {

    // 循环版本
    public static int binarySearch(int[] nums, int val) {
        if (nums == null || nums.length == 0) {
            return -1;
        }
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (nums[mid] == val) {
                return mid;
            } else if (nums[mid] > val) {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return -1;
    }

    // 递归版本
    public static int binarySearch2(int[] nums, int val) {
        if (nums == null || nums.length == 0) {
            return -1;
        }
        return recursion(nums, 0, nums.length - 1, val);
    }
    public static int recursion(int[] nums, int lo, int hi, int val) {
        if (lo > hi) {
            return -1;
        }
        int mid = lo + ((hi - lo) >> 1);
        if (nums[mid] == val) {
            return mid;
        } else if (nums[mid] > val) {
            return recursion(nums, lo, mid - 1, val);
        } else {
            return recursion(nums, mid + 1, lo, val);
        }
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
        System.out.println(binarySearch(nums, 5));
        System.out.println(binarySearch2(nums, 5));
        System.out.println(binarySearch(nums, 10));
        System.out.println(binarySearch2(nums, 0));
    }
}

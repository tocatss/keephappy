package binsearch

// Although the basic idea of binary search is comparatively straightforward,
// the details can be surprisingly tricky...
// 二分法思路很简单，细节是魔鬼
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/er-fen-cha-zhao-suan-fa-xi-jie-xiang-jie-by-labula/

// 大数字相加防止mid溢出:
// mid=(left+right)/2 => mid=left+(right-left)/2
// 循环终止条件:
// 1. left = 0,right = len-1, 为闭区间 [left,right]
// 这时终止条件是 while(left <= right) 即 [right+1,right]
// 2. left =0, right = len, 为左闭右开区间 [left,right)
// 这时终止条件是 while(left < right) 即 [right,right)

// SimpleSearch 简单二分.可以找到target但无法找到左右边界
func SimpleSearch(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		switch mid := left + (right-left)/2; {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			right = mid - 1
		case nums[mid] > target:
			left = mid + 1
		}
	}
	return -1
}

func LeftSearch(nums []int, target int) int {
	left, right := 0, len(nums) // 左闭右开区间
	for left < right {
		switch mid := left + (right-left)/2; {
		case nums[mid] == target:
			right = mid // 相等,接着找左边界.
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid // 右开区间,right = mid
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func RightSearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		switch mid := left + (right-left)/2; {
		case nums[mid] == target:
			left = mid + 1 // 找右边界.
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid
		}
	}

	if left == 0 || nums[left-1] != target {
		return -1
	}
	return left - 1
}

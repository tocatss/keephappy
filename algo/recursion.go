package algo

import "math"

func RecursionSum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	return data[0] + RecursionSum(data[1:])
}

func RecursionMax(data []int) int {
	if len(data) == 0 {
		return math.MinInt64
	}

	if len(data) == 1 {
		return data[0]
	}

	i := data[0]
	j := RecursionMax(data[1:])
	if i > j {
		return i
	}
	return j
}

// return index of target.
func RecursionBinSearch(data []int, target int) int {
	if len(data) == 0 {
		return math.MinInt64
	}
	if len(data) == 1 && data[0] == target {
		return 0
	} else if len(data) == 1 {
		return math.MinInt64
	}

	left := 0
	right := len(data)
	mid := left + (right-left)/2
	if data[mid] > target {
		right = mid
	} else if data[mid] == target {
		return mid
	} else {
		left = mid + 1
	}

	return left + RecursionBinSearch(data[left:right], target)
}

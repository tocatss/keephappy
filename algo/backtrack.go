// Package: tracereverse is some practice from
// https://liweiwei1419.gitee.io/leetcode-algo/leetcode-by-tag/backtracking/0046-permutations.html
package algo

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type intStack []int

func (s *intStack) push(item int) {
	*s = append(*s, item)
}
func (s *intStack) pop() int {
	ss := *s
	if len(ss) == 0 {
		return 0
	}
	*s = ss[0 : len(ss)-1]
	return ss[len(ss)-1]
}
func (s *intStack) equal(target *intStack) bool {
	return reflect.DeepEqual(*s, *target)
}
func (s *intStack) len() int {
	return len(*s)
}
func (s *intStack) copy() *intStack {
	copied := make(intStack, len(*s))
	copy(copied, *s)
	return &copied
}
func (s *intStack) sum() int {
	sum := 0
	for _, v := range *s {
		sum += v
	}
	return sum
}

type intSortable []int

func (s intSortable) Len() int           { return len(s) }
func (s intSortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intSortable) Less(i, j int) bool { return s[i] < s[j] }

func Permute(nums []int) [][]int {
	var (
		res   [][]int
		dfs   func(visited map[int]interface{}, stack *intStack)
		stack = make(intStack, 0)
	)
	dfs = func(visited map[int]interface{}, stack *intStack) {
		if len(*stack) == len(nums) {
			copied := make(intStack, len(*stack))
			copy(copied, *stack)
			res = append(res, copied)
			return
		}

		for _, v := range nums {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				stack.push(v)
				dfs(visited, stack)
				delete(visited, v)
				_ = stack.pop()
			}
		}
	}

	visited := make(map[int]interface{})
	dfs(visited, &stack)

	return res
}

func PermuteUnique(nums []int) [][]int {
	var (
		ans     [][]int
		dfs     func(visited []bool, path *intStack)
		visited = make([]bool, len(nums))
		path    = make(intStack, 0, len(nums))
	)

	sort.Sort(intSortable(nums))
	isRepeated := func(visited []bool, i int) bool {
		if i <= 0 {
			return false
		}
		if visited[i-1] {
			return false
		}
		return nums[i] == nums[i-1]
	}
	dfs = func(visited []bool, path *intStack) {
		if path.len() == len(nums) {
			c := path.copy()
			ans = append(ans, *c)
			return
		}
		// lastVistited := -10000
		for i, v := range nums {
			if visited[i] {
				continue
			}
			// if nums[i] == lastVistited {
			// 	continue
			// }
			if isRepeated(visited, i) {
				continue
			}
			path.push(v)
			visited[i] = true

			dfs(visited, path)
			visited[i] = false
			_ = path.pop()
			// lastVistited = nums[i]
		}
	}

	dfs(visited, &path)

	return ans
}

func CombinationSum(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		dfs  func(s *intStack, index int) bool
		path = make(intStack, 0, len(candidates))
	)

	sort.Sort(intSortable(candidates))

	dfs = func(path *intStack, index int) bool {
		if path.sum() == target {
			c := path.copy()
			ans = append(ans, *c)
			// break outer loop.
			return true
		}
		if path.sum() > target {
			// break outer loop.
			return true
		}

		for i := index; i < len(candidates); i++ {
			path.push(candidates[i])
			isBreak := dfs(path, i)
			_ = path.pop()
			if isBreak {
				// only break this loop
				return false
			}
		}
		return false
	}

	dfs(&path, 0)

	return ans
}

func CombinationSum2(candidates []int, target int) [][]int {
	var (
		ans  = make([][]int, 0)
		path = make(intStack, 0, len(candidates))
		dfs  func(path *intStack, index int) bool
	)

	dfs = func(path *intStack, index int) bool {
		sum := path.sum()
		if sum == target {
			c := path.copy()
			ans = append(ans, *c)
			return true
		}
		if sum > target {
			return true
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			path.push(candidates[i])
			isBreak := dfs(path, i+1)
			_ = path.pop()
			if isBreak {
				return false
			}
		}
		return false
	}
	sort.Sort(intSortable(candidates))
	_ = dfs(&path, 0)
	return ans
}

// 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
// 给定 n 和 k，返回第 k 个排列。
// 示例 1:
// 输入: n = 3, k = 3 (123,132,213)
// 输出: "213"
func GetPermutation(n int, k int) string {
	if n <= 0 || k <= 0 {
		return ""
	}

	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	var (
		res     string
		times   int
		visited = make([]bool, len(nums))
		dfs     func(path *stack, visited []bool) int
	)
	dfs = func(path *stack, visited []bool) int {
		if path.len() == len(nums) {
			res = path.toString()
			times++
			return times
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			path.push(nums[i])
			visited[i] = true
			if t := dfs(path, visited); t == k {
				return k
			}
			path.pop()
			visited[i] = false
		}
		return times
	}

	dfs(&stack{}, visited)

	return res
}

func GenerateParenthesis(n int) []string {
	var (
		ans     = make([]string, 0)
		visited = make([]bool, n*2)
		data    = make([]string, n*2)
		path    string
		dfs     func(path string, visited []bool, left, right int)
	)
	for i := 0; i < n; i++ {
		data[i] = "("
		data[n*2-i-1] = ")"
	}

	dfs = func(path string, visited []bool, left, right int) {
		if len(path) == n*2 {
			ans = append(ans, path)
			return
		}
		lastVisited := "-"
		for i, v := range data {
			if visited[i] {
				continue
			}
			// Skip same item.
			if v == lastVisited {
				continue
			}
			// Skip illegal item.
			if v == ")" {
				if left < right+1 {
					continue
				}
				right++
			} else {
				left++
			}

			lastVisited = v
			visited[i] = true
			path = path + v
			dfs(path, visited, left, right)

			// Back.
			if v == ")" {
				right--
			} else {
				left--
			}
			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(path, visited, 0, 0)

	return ans

}

// 给定一个无序的整数数组，找到其中上升子序列的组合

// 示例:

// 输入: [10,9,2,5,3,7,101,18]
// 输出: [10,101],[10,18] [9,101] [9,18] ...
// func findIncrease(nums []int) [][]int {
// 	ans := [][]int{}
// 	dfs := func(index int, result []int)
// 	dfs = func (index int, result []int) {

// 		for i:=index+1; i<len(nums); i++ {
// 			if nums[i] > nums[index] {
// 				// push
// 				result = append(result,nums[i])
// 				dfs(i,result)
// 				// pop
// 				result = result[:len(result)-1]
// 			}
// 		}

// 	}

// }

type stack struct {
	buff []int
}

func (s *stack) pop() {
	if len(s.buff) == 0 {
		return
	}
	s.buff = s.buff[:len(s.buff)-1]
}

func (s *stack) push(num int) {
	s.buff = append(s.buff, num)
}

func (s *stack) len() int {
	return len(s.buff)
}

func (s *stack) copy() []int {
	dst := make([]int, len(s.buff))
	copy(dst, s.buff)
	return dst
}

func (s *stack) toString() string {
	var ss strings.Builder
	for _, v := range s.buff {
		ss.WriteString(strconv.Itoa(v))
	}
	return ss.String()
}

func CombinationSum3(candidates []int, target int) [][]int {
	var (
		dfs func(path *stack, sum, from int) bool
		ans [][]int
	)

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	dfs = func(path *stack, sum, from int) bool {
		if sum == target {
			ans = append(ans, path.copy())
			return true
		}
		if sum > target {
			return true
		}

		for i := from; i < len(candidates); i++ {
			path.push(candidates[i])
			isBreak := dfs(path, sum+candidates[i], i)
			path.pop()
			if isBreak {
				return false
			}
		}
		return false
	}
	dfs(&stack{}, 0, 0)
	return ans
}

// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
// 返回 s 所有可能的分割方案。
// 示例:
// 输入: "aab"
// 输出:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]
func partition(s string) [][]string {
	// 截取a, 截取aa, 截取aab到空。
	var (
		isCircle = func(s string) bool {
			var sb strings.Builder
			for i := len(s) - 1; i >= 0; i-- {
				sb.WriteByte(s[i])
			}
			return s == sb.String()
		}
		ans [][]string
		dfs func(c *cutter, base string)
	)
	dfs = func(c *cutter, base string) {
		if base == "" {
			ans = append(ans, c.copy())
			return
		}

		for i := 0; i < len(base); i++ {
			cutted, left := c.cut(base, i)
			if !isCircle(cutted) {
				continue
			}
			c.push(cutted)
			dfs(c, left)
			c.pop()
		}
	}

	dfs(&cutter{}, s)
	return ans
}

type cutter struct {
	data []string
	// cutted string
}

// from:0 return cutted,left
func (c *cutter) cut(base string, to int) (string, string) {
	if to+1 == len(base) {
		return base, ""
	}

	return base[:to+1], base[to+1:]
}

//
func (c *cutter) push(cutted string) {
	c.data = append(c.data, cutted)
}

func (c *cutter) pop() {
	if len(c.data) == 0 {
		return
	}
	c.data = c.data[:len(c.data)-1]
}

func (c *cutter) copy() []string {
	buff := make([]string, len(c.data))
	copy(buff, c.data)
	return buff
}

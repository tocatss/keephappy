// Package: tracereverse is some practice from
// https://liweiwei1419.gitee.io/leetcode-algo/leetcode-by-tag/backtracking/0046-permutations.html
package tracereverse

import (
	"reflect"
	"sort"
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

	dfs = func(visited []bool, path *intStack) {
		if path.len() == len(nums) {
			c := path.copy()
			ans = append(ans, *c)
			return
		}
		copiedPath := path.copy()
		for i, v := range nums {
			if visited[i] {
				continue
			}
			path.push(v)
			if path.equal(copiedPath) {
				_ = path.pop()
				continue
			}
			visited[i] = true

			dfs(visited, path)

			copiedPath = path.copy()
			visited[i] = false
			_ = path.pop()
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

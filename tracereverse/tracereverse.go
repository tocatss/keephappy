// Package: tracereverse is some practice from
// https://liweiwei1419.gitee.io/leetcode-algo/leetcode-by-tag/backtracking/0046-permutations.html
package tracereverse

import (
	"log"
	"reflect"
)

func Permute(nums []int) [][]int {
	stack := make(intStack, 0)
	var (
		res [][]int
		dfs func(visited map[int]bool, stack *intStack)
	)
	dfs = func(visited map[int]bool, stack *intStack) {
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

	visited := make(map[int]bool)
	for _, v := range nums {
		visited[v] = true
		stack.push(v)
		dfs(visited, &stack)
		delete(visited, v)
		_ = stack.pop()
	}

	return res
}

// TODO: not complete.
func PermuteUnique(nums []int) [][]int {
	var (
		ans      [][]int
		dfs      func(visisted []bool, path *intStack)
		visisted = make([]bool, len(nums))
	)

	dfs = func(visisted []bool, path *intStack) {
		if path.len() == len(nums) {
			c := path.copy()
			ans = append(ans, *c)
			return
		}
		copiedPath := path.copy()
		for i, v := range nums {
			if visisted[i] {
				continue
			}
			path.push(v)
			if path.equal(copiedPath) {
				_ = path.pop()
				continue
			}
			visisted[i] = true

			dfs(visisted, path)

			copiedPath = path.copy()
			visisted[i] = false
			_ = path.pop()
		}
	}

	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		path := make(intStack, 0, len(nums))
		visisted[i] = true
		path.push(v)

		dfs(visisted, &path)

		visisted[i] = false
		path.pop()
	}
	log.Print(ans)
	return nil
}

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

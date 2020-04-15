// https://liweiwei1419.gitee.io/leetcode-algo/leetcode-by-tag/backtracking/0046-permutations.html
package tracereverse

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

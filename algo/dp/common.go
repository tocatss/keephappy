// https://liweiwei1419.gitee.io/leetcode-algo/leetcode-by-tag/dynamic-programming/
package dp

import (
	"fmt"
	"math"
)

// dynamic programming 𝑎𝑙𝑠𝑜𝑘𝑛𝑜𝑤𝑛𝑎𝑠𝑑𝑦𝑛𝑎𝑚𝑖𝑐𝑜𝑝𝑡𝑖𝑚𝑖𝑧𝑎𝑡𝑖𝑜𝑛 is a method for solving a complex problem
// by breaking it down into a collection of simpler subproblems, solving each of those subproblems just once,
// and storing their solutions – ideally, using a memory-based data structure.

// 对于一个递归结构的问题，如果我们在分析它的过程中，发现了它有很多“重叠子问题” => 重复运算，
// 虽然并不影响结果的正确性，但是我们认为大量的重复计算是不简洁，不优雅，不高效的，因此，我们必须将“重叠子问题”进行优化，
// 优化的方法就是“加入缓存”，“加入缓存”的一个学术上的叫法就是“记忆化搜索”。
// 另外，我们还发现，直接分析递归结构，是假设更小的子问题已经解决给出的实现，思考的路径是“自顶向下”。
// 但有的时候，“自底向上”的思考路径往往更直接，这就是“动态规划”，
// 我们是真正地解决了更小规模的问题，
// 在处理更大规模的问题的时候，直接使用了更小规模问题的结果。

// Fib: f(0) = 0,f(1) = 1 => f(n) = f(n-1)+f(n-2)
func fibrecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibrecursion(n-1) + fibrecursion(n-2)
}

func fibdynamic(n int) int {
	memo := make([]int, 0, n+1)
	memo = append(memo, 0, 1)

	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[n]
}

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 递归:仍然存在重复计算.
// n阶 = n-1 阶 + n-2 阶 => fib
// https://leetcode-cn.com/problems/climbing-stairs/description/
func ClimbingStairsRecursion(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2 // [1,1],[2]
	}

	return ClimbingStairsRecursion(n-1) + ClimbingStairsRecursion(n-2)
}

// 记忆化搜索版: 由上向下递归并缓存
func ClimbingStairsRecWithMemo(n int) int {
	if n < 0 {
		return 0
	}

	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2

	var climb func(n int) int
	climb = func(k int) int {
		if v, ok := memo[k]; ok {
			return v
		}
		memo[k] = climb(k-1) + climb(k-2)
		return memo[k]
	}

	return climb(n)
}

// 动态规划: 自底向上.
func ClimbingStairsDynamic(n int) int {
	if n <= 0 {
		return 0
	}

	memo := make([]int, 0, n)
	memo = append(memo, 1, 2)

	for i := 2; i < n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[n-1]
}

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

// 动态规划：
// 动态规划告诉我们解决问题的思路，我们不是直接去解决题目问的问题，而是去发现这个问题最开始的样子，
// 通过「状态转移」，每一步参考了之前计算的结果，得到最终的答案。
// 解：
// dp[i] => 以nums[i]为终点的最大和
// dp[i-1] => 以nums[i-1]为终点的最大和
// dp[i] = Math.Max(nums[i], dp[i-1] + nums[i])  // dp[i-1] 为可能为负
// 本题就是要找到 dp[0]...dp[n-1]的最大值。
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	ans := nums[0]

	// 由于下一个状态只和上一个有关，那么可以不用数组，从而缩小空间复杂度
	// dp := make([]int, len(nums))
	// dp[0] = nums[0]
	state := nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		if state < 0 {
			state = nums[i]
			if state > ans {
				ans = state
			}
			continue
		}

		state += nums[i]
		if state > ans {
			ans = state
		}
	}

	return ans
}

// 输入: [10,9,2,5,3,7,101,18]
// 输出: 4
// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[n] 到第N个元素的个数
	dp := make([]int, len(nums))
	// 找到比index小的元素所拥有最长的上升子序列个数的最大值。
	maxOfLower := func(index int) int {
		max := 0
		target := nums[index]
		for i := index - 1; i >= 0; i-- {
			if nums[i] < target && dp[i] > max {
				max = dp[i]
			}
		}
		return max
	}

	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxOfLower(i) + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 输入: m = 3, n = 2
// 输出: 3
// 解释:
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向右 -> 向下
// 2. 向右 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向右
func UniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return -1
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// start dp[0][0]
	// end   dp[n-1][m-1]
	for i := range dp {
		for j := range dp[i] {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}

	}

	return dp[n-1][m-1]
}

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return -1
	}
	if len(obstacleGrid[0]) == 0 {
		return -1
	}

	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := range obstacleGrid {
		for j, v := range obstacleGrid[i] {
			if v == 1 {
				dp[i][j] = 0
				continue
			}

			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[0][j] = dp[0][j-1]
			} else if j == 0 {
				dp[i][0] = dp[i-1][0]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[n-1][m-1]
}

// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 7
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])

	to := make([][]int, n)
	for i := 0; i < n; i++ {
		to[i] = make([]int, m)
		copy(to[i], grid[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				to[i][j] += to[i][j-1]
			} else if j == 0 {
				to[i][j] += to[i-1][j]
			} else {
				if to[i-1][j] < to[i][j-1] {
					to[i][j] += to[i-1][j]
				} else {
					to[i][j] += to[i][j-1]
				}
			}
		}
	}
	return to[n-1][m-1]
}

// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符
//

// 示例 1：

// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')

// s1 => s2的转换方法：一般都是用两个指针 i,j 分别指向两个字符串的最后，然后一步步往前走，缩小问题的规模。
// https://leetcode-cn.com/circle/article/rY3PIQ/
func minDistanceRecursion(word1 string, word2 string) int {
	var (
		// dp: 返回word1[:i] 和 word2[:j]的最少操作数。
		dp func(i, j int) int
		// min: 返回i,j,k中的最小值。
		min func(i, j, k int) int = func(i, j, k int) int {
			min := i
			if j < min {
				min = j
			}
			if k < min {
				min = k
			}
			return min
		}
	)

	dp = func(i, j int) int {
		// word1 的指针移动到头了。
		if i == -1 {
			// word2剩下的都插入就好了。
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		// 相等：word1[:i] 和 word2[:j]的最少操作数 就等同于 word1[:i-1] 和 word2[:j-1]的最少操作数.
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		}

		// 不等，那么我们可以选择 增加，或者删除，或者替换。
		return min(
			// 向word1的末尾增加word2[j], 所以j向前。
			dp(i, j-1)+1,
			// 删除word1的元素，所以i向前。
			dp(i-1, j)+1,
			// word1[i] 替换成word2[j]
			dp(i-1, j-1)+1,
		)
	}

	return dp(len(word1)-1, len(word2)-1)
}

// 由上例可以发现
func minDistanceRecursionWithMemo(word1 string, word2 string) int {
	var (
		dp  func(i, j int) int
		min = func(i, j, k int) int {
			min := i
			if j < min {
				min = j
			}
			if k < min {
				min = k
			}
			return min
		}
		memo   = map[string]int{}
		genKey = func(i, j int) string {
			return fmt.Sprintf("%d-%d", i, j)
		}
	)

	dp = func(i, j int) int {
		k := genKey(i, j)
		if v, ok := memo[k]; ok {
			return v
		}

		if i == -1 {
			memo[k] = j + 1
			return j + 1
		}
		if j == -1 {
			memo[k] = i + 1
			return i + 1
		}

		if word1[i] == word2[j] {
			v := dp(i-1, j-1)
			memo[k] = v
			return v
		}

		v := min(
			dp(i, j-1)+1,
			dp(i-1, j)+1,
			dp(i-1, j-1)+1,
		)
		memo[k] = v
		return v
	}
	return dp(len(word1)-1, len(word2)-1)
}

func minDistanceDP(word1 string, word2 string) int {
	min := func(i, j, k int) int {
		min := i
		if j < min {
			min = j
		}
		if k < min {
			min = k
		}
		return min
	}
	// dp 是二维数组用来递进推出 dp[i][j]
	// dp[i][j] =  min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1,)

	// 初始化
	dp := make([][]int, len(word1)+1)

	// dp[0][1] ...dp[0][j] => "" ... word2
	// word1 同理
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
		for j := 0; j < len(word2)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
					continue
				}
				dp[i][j] = min(
					dp[i-1][j-1]+1,
					dp[i][j-1]+1,
					dp[i-1][j]+1,
				)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
// 示例 1：

// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：

// 输入: "cbbd"
// 输出: "bb"
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// dp[i][j]: s[i...j]闭区间 是否为回文。
	// j > i
	dp := make([][]bool, len(s))
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if j == i {
				dp[i][j] = true
			} else if j <= i+2 {
				dp[i][j] = s[j] == s[i]
			} else {
				if s[j] != s[i] {
					dp[i][j] = false
				}
				dp[i][j] = dp[i+1][j-1]
			}
		}
	}
	return ""
}

// No test case because longestCommonSubsequence has passed in leetcode.
// https://leetcode-cn.com/problems/longest-common-subsequence/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace"，它的长度为 3。

// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0。
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	var (
		findMax = func(p, q int) int {
			if p > q {
				return p
			}
			return q
		}
		// dp[i][j]表示text1[:i+1]和text2[:j+1]中最长的公共子序列
		dp [][]int
	)
	dp = make([][]int, len(text1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2))
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if i == 0 && j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
					continue
				}
				dp[0][0] = 0
			} else if i == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
					continue
				}
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
					continue
				}
				dp[i][j] = dp[i-1][j]
			} else {
				if text1[i] == text2[j] {
					dp[i][j] = dp[i-1][j-1] + 1
					continue
				}
				dp[i][j] = findMax(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}

// Rob...
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
// 示例 1：
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。
// https://leetcode-cn.com/problems/house-robber/
// Funtion has passed in leetcode.
func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	// dp[i]: 到i家可以偷到的最多的money
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		if dp[i-1] > dp[i-2]+nums[i] {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = dp[i-2] + nums[i]
	}
	return dp[len(nums)-1]
}

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
// 示例 1:
// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3
// 解释: 11 = 5 + 5 + 1
//
// 示例 2:
// 输入: coins = [2], amount = 3
// 输出: -1
// Funtion has passed at leetcode.
// https://leetcode-cn.com/problems/coin-change/submissions/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// dp[amount] 到amount为止最少的组合次数。
	// 如果dp[amount-coins[1] 存在
	// dp[amount] = min(dp[amount-coins[1]] + 1, dp[amount-coins[2]] + 1,...)
	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		min := math.MaxInt64
		for i := 0; i < len(coins); i++ {
			if coins[i] > a {
				continue
			}
			if dp[a-coins[i]] == -1 {
				continue
			}
			// find min.
			if dp[a-coins[i]]+1 < min {
				min = dp[a-coins[i]] + 1
			}
		}
		// not found
		if min == math.MaxInt64 {
			dp[a] = -1
			continue
		}

		dp[a] = min
	}

	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

// // 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。
// // 假设每一种面额的硬币有无限个。
// // 示例 1:
// // 输入: amount = 5, coins = [1, 2, 5]
// // 输出: 4
// // 解释: 有四种方式可以凑成总金额:
// // 5=5
// // 5=2+2+1
// // 5=2+1+1+1
// // 5=1+1+1+1+1

// // 示例 2:
// // 输入: amount = 3, coins = [2]
// // 输出: 0
// // 解释: 只用面额2的硬币不能凑成总金额3。
// // https://leetcode-cn.com/problems/coin-change-2
// func change(amount int, coins []int) int {
// 	if amount == 0 {
// 		return 0
// 	}

// 	// dp[i]: 凑成总值为i的个数
// 	dp := make([]int, amount+1)
// 	for i := 1; i < amount; i++ {
// 		for j := 0; j < len(coins); j++ {
// 			if coins[j] > i {
// 				continue
// 			}

// 		}
// 	}
// }

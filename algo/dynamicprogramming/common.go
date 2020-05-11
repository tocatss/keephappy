package dynamicprogramming

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
	// 找到比index小的最大值所拥有的元素个数。
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

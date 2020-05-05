package dynamicopt

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

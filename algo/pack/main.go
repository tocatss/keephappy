// 学习一下背包问题九讲 2.0 beta1.2(Tianyi Cui)
// https://github.com/tianyicui/pack/blob/master/V2.pdf

package main

// 问题：有 N 件物品和一个容量为 V 的背包。
//      放入第 i 件物品耗费的费用是 Ci1，得到的 价值是 Wi。求解将哪些物品装入背包可使价值总和最大。
// 最基础的背包问题，特点是:每种物品仅有一件，可以选择放或不放。
// 用子问题定义状态:即 F [i, v] 表示前 i 件物品恰放入一个容量为 v 的背包可以获得 的最大价值。则其状态转移方程便是:
// F[i,v] = max{F[i − 1,v],F[i − 1,v − Ci] + Wi}

// vs: value map
// cs: cost map
// max: caps of pack
func ZeroOnePack(max int, vs, cs map[string]int) int {
	if max <= 0 || len(vs) == 0 {
		return 0
	}

	ik := make(map[int]string, len(vs))
	dp := make([][]int, len(vs)+1)

	// init
	dp[0] = make([]int, max+1)
	i := 1
	for k := range vs {
		dp[i] = make([]int, max+1) // 0...max
		ik[i] = k
		i++
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j <= max; j++ {
			k := ik[i]
			cost := cs[k]
			value := vs[k]
			if j < cost {
				dp[i][j] = dp[i-1][j]
				continue
			}

			before := dp[i-1][j]
			after := dp[i-1][j-cost] + value
			if before > after {
				dp[i][j] = before
			} else {
				dp[i][j] = after
			}
		}
	}

	return dp[len(vs)][max]
}

func ZeroOnePackOptimization(max int, vs, cs map[string]int) int {
	old := make([]int, max+1)
	for k, v := range vs {
		cost := cs[k]

		new := make([]int, max+1)
		copy(new, old)
		for i := cost; i <= max; i++ {
			before := old[i]
			after := old[i-cost] + v
			if before < after {
				new[i] = after
			}
		}
		old = new
	}
	return old[max]

	// 另一种采用递减的方式仍可以实现
	// dp := make([]int, max+1)
	// for k, v := range vs {
	// 	cost := cs[k]
	// 	for i := max; i >= cost; i-- {
	// 		before := dp[i]
	// 		after := dp[i-cost] + v
	// 		if before < after {
	// 			dp[i] = after
	// 		}
	// 	}
	// }
	// return dp[max]
}

// 完全背包问题
// 问题 有 N 种物品和一个容量为 V 的背包，每种物品都有无限件可用。
//     放入第 i 种物品 的费用是 Ci，价值是 Wi。求解:将哪些物品装入背包，可使这些物品的耗费的费用总 和不超过背包容量，且价值总和最大

// F[i,v]=max{F[i−1,v−kCi]+kWi |0 ≤ kCi ≤ v}
func CompletePack(max int, vs, cs map[string]int) int {
	if max <= 0 || len(vs) == 0 {
		return 0
	}

	dp := make([][]int, len(vs)+1)

	// init
	dp[0] = make([]int, max+1)
	i := 1
	ik := make(map[int]string, len(vs))
	for k := range vs {
		dp[i] = make([]int, max+1)
		ik[i] = k
		i++
	}

	for i := 1; i < len(dp); i++ {
		k := ik[i]
		cost := cs[k]
		value := vs[k]
		for j := 1; j <= max; j++ {
			m := 0
			for n := 0; n*cost <= j; n++ {
				nv := dp[i-1][j-n*cost] + n*value
				if m < nv {
					m = nv
				}
			}
			dp[i][j] = m
		}
	}

	return dp[len(vs)][max]
}

func CompletePackOptimization(v int, ws, cs map[string]int) int {
	dp := make([]int, v+1)
	for k, cost := range cs {
		for i := cost; i < len(dp); i++ {
			w := ws[k]
			max := dp[i]
			for j := 0; j*cost <= i; j++ {
				if dp[i-j*cost]+w*j > max {
					max = dp[i-j*cost] + w*j
				}
			}
			dp[i] = max
		}
	}

	return dp[v]
	// 01 背包中要按照 v 递减的次序来循环。 让 v 递减是为了保证第 i 次循环中的状态 F[i,v] 是由状态 F[i − 1,v − Ci] 递推而来。
	// 换句话说，这正是为了保证每件物品只选一次，保证在考虑“选入第 i 件物品”这件策 略时，依据的是一个绝无已经选入第 i 件物品的子结果 F [i − 1, v − Ci]。
	// 而现在完全背 包的特点恰是每种物品可选无限件，所以在考虑“加选一件第 i 种物品”这种策略时， 却正需要一个可能已选入第 i 种物品的子结果 F [i, v − Ci ]
	// 所以就可以并且必须采用 v 递增的顺序循环。这就是这个简单的程序为何成立的道理。

	// 递增 dp[i] 代表容量为i时当前的最大价值。
	// dp := make([]int, max+1)
	// for k, v := range vs {
	// 	cost := cs[k]
	// 	for i := cost; i <= max; i++ {
	// 		after := dp[i-cost] + v
	// 		if dp[i] < after {
	// 			dp[i] = after
	// 		}
	// 	}
	// }
	// return dp[max]

	// 递减
	// dp := make([]int, max+1)
	// for k := range vs {
	// 	cost := cs[k]
	// 	value := vs[k]

	// 	for i := max; i >= cost; i-- {
	// 		m := 0
	// 		for n := 0; n*cost <= i; n++ {
	// 			nv := dp[i-n*cost] + n*value
	// 			if nv > m {
	// 				m = nv
	// 			}
	// 		}
	// 		dp[i] = m
	// 	}
	// }
	// return dp[max]
}

// 多重背包问题
// 问题 有 N 种物品和一个容量为 V 的背包。第 i 种物品最多有 Mi 件可用，每件耗费的 空间是 Ci，价值是 Wi。
//     求解将哪些物品装入背包可使这些物品的耗费的空间总和不超 过背包容量，且价值总和最大。

// 这题目和完全背包问题很类似。基本的方程只需将完全背包问题的方程略微一改 即可。
// 因为对于第 i 种物品有 Mi +1 种策略:取 0 件，取 1 件......取 Mi 件。令 F[i,v] 表示前 i 种物品恰放入一个容量为 v 的背包的最大价值，则有状态转移方程:
// F [i，v] = max{F [i − 1, v − k ∗ Ci] + k ∗ Wi | 0 ≤ k ≤ Mi} 复杂度是 O(V ΣMi)。

func MultiplePack(max int, vs, cs, ns map[string]int) int {
	if max <= 0 || len(vs) == 0 {
		return 0
	}

	dp := make([][]int, len(vs)+1)

	// init
	dp[0] = make([]int, max+1)
	i := 1
	ik := make(map[int]string, len(vs))
	for k := range vs {
		dp[i] = make([]int, max+1)
		ik[i] = k
		i++
	}

	for i := 1; i < len(dp); i++ {
		k := ik[i]

		cost := cs[k]
		value := vs[k]
		number := ns[k] // 最大件数
		for j := 1; j <= max; j++ {
			m := dp[i-1][j]
			for n := 0; n*cost <= j && n <= number; n++ {
				after := dp[i-1][j-n*cost] + n*value
				if after > m {
					m = after
				}
			}
			dp[i][j] = m
		}
	}

	return dp[len(vs)][max]
}

func MultiplePackOptimization(max int, vs, cs, ns map[string]int) int {
	// 递减
	dp := make([]int, max+1)

	for k, v := range vs {
		cost := cs[k]
		number := ns[k]

		for i := max; i >= cost; i-- {
			m := 0
			for n := 0; n*cost <= i && n <= number; n++ {
				after := dp[i-n*cost] + n*v
				if after > m {
					m = after
				}
			}
			dp[i] = m
		}
	}

	return dp[max]
}

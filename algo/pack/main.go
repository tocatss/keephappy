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
	if max == 0 || len(vs) == 0 {
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
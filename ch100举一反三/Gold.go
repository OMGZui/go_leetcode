/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/16
 * Time: 09:48
 */
package main

/* 有坑，回填 */

// 5座金矿   10个工人    取得最大金矿
// n/金矿数量 w/工人数量 g/金矿存量 p/金矿需要工人数
func getMuchGold(n int, w int, g []int, p []int) int {
	preRes := make([]int, len(p))
	res := make([]int, len(p))

	// 填充边界的格子值
	for i := 0; i <= w; i++ {
		// 工人数小于金矿需要工人数，直接取0
		if i < p[0] {
			preRes[i] = 0
		} else {
			preRes[i] = g[0]
		}
	}

	// 填充其余格子
	// 外层循环金矿数量
	for i := 0; i < n; i++ {
		// 内层循环工人数量
		for j := 0; j <= w; j++ {
			if j < p[i] {
				res[j] = preRes[j]
			} else {
				res[j] = max(preRes[j], preRes[j-p[i]]+g[i])
			}
		}
		preRes = res
	}
	return preRes[w]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	//n, w := 5, 10
	//g := []int{400, 500, 200, 300, 350}
	//p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(getMuchGold(n, w, g, p))
}

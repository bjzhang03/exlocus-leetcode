package medium_0688

func knightProbability(N int, K int, r int, c int) float64 {
	// k==0 的时候,此时不需要移动,已经在棋盘内部了,不会再移动了
	if N > 0 && K > 0 {
		return knightMove(map[int]float64{pointToNum([]int{r, c}): 1}, N, K)
	}
	return 1
}

// 使用浮点数存在精度损失的问题
func knightMove(points map[int]float64, N int, step int) float64 {
	if step > 0 && len(points) > 0 {
		currentsum := float64(0)
		npoints := map[int]float64{}
		for key, val := range points {
			currentsum = currentsum + float64(val)
			point := keyToPoint(key)
			if pass(N, []int{point[0] - 1, point[1] - 2}) {
				if _, ok := npoints[pointToNum([]int{point[0] - 1, point[1] - 2})]; ok {
					npoints[pointToNum([]int{point[0] - 1, point[1] - 2})] = npoints[pointToNum([]int{point[0] - 1, point[1] - 2})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] - 1, point[1] - 2})] = + val
				}
			}

			if pass(N, []int{point[0] - 1, point[1] + 2}) {
				if _, ok := npoints[pointToNum([]int{point[0] - 1, point[1] + 2})]; ok {
					npoints[pointToNum([]int{point[0] - 1, point[1] + 2})] = npoints[pointToNum([]int{point[0] - 1, point[1] + 2})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] - 1, point[1] + 2})] = + val
				}
			}

			if pass(N, []int{point[0] + 1, point[1] - 2}) {
				if _, ok := npoints[pointToNum([]int{point[0] + 1, point[1] - 2})]; ok {
					npoints[pointToNum([]int{point[0] + 1, point[1] - 2})] = npoints[pointToNum([]int{point[0] + 1, point[1] - 2})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] + 1, point[1] - 2})] = + val
				}
			}
			if pass(N, []int{point[0] + 1, point[1] + 2}) {
				if _, ok := npoints[pointToNum([]int{point[0] + 1, point[1] + 2})]; ok {
					npoints[pointToNum([]int{point[0] + 1, point[1] + 2})] = npoints[pointToNum([]int{point[0] + 1, point[1] + 2})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] + 1, point[1] + 2})] = + val
				}
			}
			if pass(N, []int{point[0] - 2, point[1] - 1}) {
				if _, ok := npoints[pointToNum([]int{point[0] - 2, point[1] - 1})]; ok {
					npoints[pointToNum([]int{point[0] - 2, point[1] - 1})] = npoints[pointToNum([]int{point[0] - 2, point[1] - 1})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] - 2, point[1] - 1})] = + val
				}
			}
			if pass(N, []int{point[0] - 2, point[1] + 1}) {
				if _, ok := npoints[pointToNum([]int{point[0] - 2, point[1] + 1})]; ok {
					npoints[pointToNum([]int{point[0] - 2, point[1] + 1})] = npoints[pointToNum([]int{point[0] - 2, point[1] + 1})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] - 2, point[1] + 1})] = + val
				}
			}

			if pass(N, []int{point[0] + 2, point[1] - 1}) {
				if _, ok := npoints[pointToNum([]int{point[0] + 2, point[1] - 1})]; ok {
					npoints[pointToNum([]int{point[0] + 2, point[1] - 1})] = npoints[pointToNum([]int{point[0] + 2, point[1] - 1})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] + 2, point[1] - 1})] = + val
				}
			}
			if pass(N, []int{point[0] + 2, point[1] + 1}) {
				if _, ok := npoints[pointToNum([]int{point[0] + 2, point[1] + 1})]; ok {
					npoints[pointToNum([]int{point[0] + 2, point[1] + 1})] = npoints[pointToNum([]int{point[0] + 2, point[1] + 1})] + val
				} else if !ok {
					npoints[pointToNum([]int{point[0] + 2, point[1] + 1})] = + val
				}
			}
		}

		// 统计一下有多少是在棋盘内部的数据
		countin := float64(0)
		for _, val := range npoints {
			countin = countin + val
		}
		return (float64(countin) / (currentsum * 8)) * knightMove(npoints, N, step-1)
	}
	return 1
}

func pass(N int, p []int) bool {
	return (p[0] >= 0 && p[0] < N && p[1] >= 0 && p[1] < N)
}

func pointToNum(p []int) int {
	return p[0] + 1000*p[1]
}

func keyToPoint(key int) []int {
	return []int{key % 1000, key / 1000}
}

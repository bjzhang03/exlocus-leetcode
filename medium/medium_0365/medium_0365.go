package medium_0365

func canMeasureWater(x int, y int, z int) bool {
	// 如果直接为0的话，可以直接返回true
	if z == 0 {
		return true
	} else if x > 0 && y > 0 && z > 0 && z <= x+y {
		// 这里有一个额外的条件,就是z不能大于x+y,这个条件我不知道如何证明
		gcd := getGCD(x, y)

		if gcd == 1 {
			return true
		} else if getGCD(gcd, z) != 1 {
			return true
		}
	}
	return false
}

// 计算最大公约数
func getGCD(x, y int) int {
	larger := x
	miner := y

	if x < y {
		larger = y
		miner = x
	}

	if larger%miner == 0 {
		return miner
	} else {
		return getGCD(miner, larger%miner)
	}
	return 0
}

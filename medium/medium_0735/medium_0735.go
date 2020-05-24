package medium_0735

import (
	"math"
)

func asteroidCollision(asteroids []int) []int {
	// 预先分配空间,优化分配机制
	nasteroids := make([]int, len(asteroids))
	for !checkNoCollision(asteroids) {
		// 标记符号
		flag := (asteroids[0] > 0)
		// 清空数据
		nasteroids = nasteroids[0:0]
		collisions := false
		for i := 0; i < len(asteroids); i++ {
			if flag != (asteroids[i] > 0) && flag {
				// 发生碰撞了
				if math.Abs(float64(asteroids[i-1])) > math.Abs(float64(asteroids[i])) {
					collisions = true
				} else if math.Abs(float64(asteroids[i-1])) < math.Abs(float64(asteroids[i])) {
					nasteroids = append(nasteroids[:len(nasteroids)-1], asteroids[i])
					collisions = true
				} else {
					nasteroids = nasteroids[:len(nasteroids)-1]
					collisions = true
				}
				// 更新符号
				if i+1 < len(asteroids) {
					flag = (asteroids[i+1] > 0)
				}
			} else {
				// 没有发生碰撞
				nasteroids = append(nasteroids, asteroids[i])
				// 更新flag的数据
				flag = (asteroids[i] > 0)
			}
		}
		// 如果发生碰撞,则更新数据
		if collisions {
			asteroids = asteroids[:0]
			asteroids = append(asteroids, nasteroids...)
		}
	}
	return asteroids
}

// 检查是否会发生碰撞
func checkNoCollision(asteroids []int) bool {
	if len(asteroids) > 0 {
		flag := (asteroids[0] > 0)
		for i := 0; i < len(asteroids); i++ {
			// 发生碰撞了
			if flag != (asteroids[i] > 0) && flag == true {
				return false
			}
			flag = (asteroids[i] > 0)
		}
	}
	return true
}

package medium_0754

// 数学题,常规方法解决不了
func reachNumber(target int) int {
	return solve(target)
}

// 这个是常规的方法
func solve(target int) int {
	visit := []int{0}
	cstep := 0
	for len(visit) > 0 {
		// 跨出了第一步
		cstep = cstep + 1
		//fmt.Printf("visit := %s, cstep:= %s\n", fmt.Sprint(visit), fmt.Sprint(cstep))
		nvisit := []int{}
		nused := map[int]bool{}
		abs := func(data int) int {
			if data < 0 {
				return -data
			}
			return data
		}
		// 挨个生成可以到达的数据
		for i := 0; i < len(visit); i++ {
			// 新生成两个数据
			if abs(visit[i]-cstep) == abs(target) {
				return cstep
			}
			if _, ok := nused[abs(visit[i]-cstep)]; !ok {
				nvisit = append(nvisit, abs(visit[i]-cstep))
				nused[abs(visit[i]-cstep)] = true
			}
			if abs(visit[i]+cstep) == abs(target) {
				return cstep
			}
			if _, ok := nused[abs(visit[i]+cstep)]; !ok {
				nvisit = append(nvisit, abs(visit[i]+cstep))
				nused[abs(visit[i]+cstep)] = true
			}
		}
		visit = append(visit[:0], nvisit...)
	}
	return 0
}

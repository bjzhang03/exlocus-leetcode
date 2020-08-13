package medium_0526

func countArrangement(N int) int {
	if N > 0 {
		// 构造数组
		save := []int{}
		for i := 1; i <= N; i++ {
			save = append(save, i)
		}
		rsave := 0
		deepFirstSearch(save, &rsave, []int{})
		return rsave
	}
	return 0
}

func deepFirstSearch(save []int, result *int, tmp []int) {
	if len(save) == 0 {
		(*result) ++
	} else if len(save) > 0 {
		for i := 0; i < len(save); i++ {
			tmp = append(tmp, save[i])
			// 剪枝操作
			if (save[i]%len(tmp) == 0 || len(tmp)%save[i] == 0) {
				// 获取下一次的待选择的数据
				deepFirstSearch(append(append([]int{}, save[:i]...), save[i+1:]...), result, tmp)
			}
			tmp = tmp[:len(tmp)-1]
		}
	}
}

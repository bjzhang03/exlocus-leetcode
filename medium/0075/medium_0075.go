package _075

func sortColors(nums []int) {

	count := make(map[int]int)

	for _, key := range nums {
		if _, ok := count[key]; ok {
			count[key]++
		} else {
			count[key] = 1
		}
	}

	indexSta := 0

	for i := 0; i <= 2; i++ {
		indexEnd := count[i]
		for j := indexSta; j < indexSta+indexEnd; j++ {
			nums[j] = i
		}

		indexSta = indexSta + indexEnd
	}
}

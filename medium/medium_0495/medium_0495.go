package medium_0495

func findPoisonedDuration(timeSeries []int, duration int) int {
	result := 0
	if len(timeSeries) > 0 {
		for i := 1; i < len(timeSeries); i++ {
			if timeSeries[i]-timeSeries[i-1] >= duration {
				result = result + duration
			} else {
				result = result + (timeSeries[i] - timeSeries[i-1])
			}
		}
		// 最后一步一定是duration
		result = result + duration
	}
	return result
}

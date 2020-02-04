package medium_0436

//You may assume the interval's end point is always bigger than its start point.
//You may assume none of these intervals have the same start point.
func findRightInterval(intervals [][]int) []int {
	result := []int{}
	if len(intervals) > 0 {
		for i := 0; i < len(intervals); i++ {
			tmpscript := -1
			tmpinterval := []int{}
			for j := 0; j < len(intervals); j++ {
				if intervals[j][0] >= intervals[i][1] {
					if len(tmpinterval) == 0 {
						tmpinterval = append([]int{}, intervals[j]...)
						tmpscript = j
					} else if tmpinterval[0] > intervals[j][0] {
						tmpinterval = append([]int{}, intervals[j]...)
						tmpscript = j
					}
				}
			}
			result = append(result, tmpscript)
		}
	}
	return result
}

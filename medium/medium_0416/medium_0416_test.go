package medium_0416

import "testing"

func TestCanPartition(t *testing.T) {

	var cp = []struct {
		nums     []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 5, 11, 3}, false},
		{[]int{28, 63, 95, 30, 39, 16, 36, 44, 37, 100, 61, 73, 32, 71, 100, 2, 37, 60, 23, 71, 53, 70, 69, 82, 97, 43, 16, 33, 29, 5, 97, 32, 29, 78, 93, 59, 37, 88, 89, 79, 75, 9, 74, 32, 81, 12, 34, 13, 16, 15, 16, 40, 90, 70, 17, 78, 54, 81, 18, 92, 75, 74, 59, 18, 66, 62, 55, 19, 2, 67, 30, 25, 64, 84, 25, 76, 98, 59, 74, 87, 5, 93, 97, 68, 20, 58, 55, 73, 74, 97, 49, 71, 42, 26, 8, 87, 99, 1, 16, 79}, true},
		{[]int{63, 12, 4, 8, 76, 12, 52, 68, 70, 59, 65, 64, 26, 45, 93, 30, 70, 62, 61, 12, 36, 90, 91, 7, 84, 56, 37, 58, 4, 64, 60, 4, 100, 66, 90, 35, 60, 6, 68, 61, 24, 41, 86, 88, 89, 90, 7, 54, 3, 20, 86, 95, 93, 77, 87, 41, 30, 13, 41, 1, 82, 59, 18, 18, 34, 53, 25, 32, 74, 82, 82, 62, 46, 21, 38, 78, 66, 67, 62, 75, 76, 17, 96, 42, 32, 44, 89, 95, 96, 93, 3, 89, 9, 92, 21, 30, 68, 38, 16, 73}, true},
		{[]int{97, 100, 88, 49, 43, 55, 2, 62, 72, 13, 60, 36, 67, 64, 13, 39, 66, 6, 26, 45, 46, 75, 28, 66, 71, 75, 91, 33, 64, 54, 3, 76, 76, 35, 49, 6, 63, 11, 80, 86, 73, 95, 60, 58, 61, 42, 52, 27, 73, 51, 58, 38, 28, 62, 84, 55, 90, 40, 52, 96, 55, 32, 52, 63, 44, 90, 14, 68, 50, 32, 73, 64, 92, 74, 66, 64, 22, 51, 27, 30, 83, 30, 37, 25, 22, 46, 95, 59, 57, 21, 29, 72, 7, 56, 47, 48, 54, 75, 67, 33}, true},
	}

	for _, val := range cp {
		actual := canPartition(val.nums)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}

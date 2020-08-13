package medium_0698

import "testing"

func TestCanPartitionKSubsets(t *testing.T) {
	var cpks = []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{960, 3787, 1951, 5450, 4813, 752, 1397, 801, 1990, 1095, 3643, 8133, 893, 5306, 8341, 5246}, 6, true},
		{[]int{7628, 3147, 7137, 2578, 7742, 2746, 4264, 7704, 9532, 9679, 8963, 3223, 2133, 7792, 5911, 3979}, 6, false},
	}

	for _, val := range cpks {
		actual := canPartitionKSubsets(val.nums, val.k)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %t, expected : %t", actual, val.expected)
		}
	}

}

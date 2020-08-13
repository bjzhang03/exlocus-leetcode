package medium_0473

import "testing"

func TestMakesquare(t *testing.T) {

	var ms = []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 1, 2, 2, 2}, true},
		{[]int{3, 3, 3, 3, 4}, false},
		{[]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}, true},
		{[]int{211559, 9514615, 7412176, 5656677, 3816020, 452925, 7979371, 5025276, 8882605, 944541, 9889007, 2344356, 7252152, 749758, 2311818}, false},
		{[]int{6122053, 1031956, 460065, 3996684, 3891947, 1839190, 6127621, 8855952, 8835594, 3425930, 4189081, 7596722, 876208, 7952011, 9676846}, false},
	}

	for _, val := range ms {
		actual := makesquare(val.nums)
		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}
}

package medium_0372

var modc = 1337

func superPow(a int, b []int) int {
	result := 1
	if a > 1 {
		a = a % modc

		save := getcount(a, b)
		for i := 0; i < len(b); i++ {
			for j := 0; j < b[i]; j++ {
				result = (result * save[i]) % modc
			}
		}

	}
	return result
}

// 临时变量,用于存储每一个的临时值
func getcount(a int, b []int) []int {
	result := make([]int, len(b))
	result[len(b)-1] = a

	for i := len(b) - 2; i >= 0; i-- {
		tmp := 1
		for j := 0; j < 10; j++ {
			tmp = (tmp * result[i+1]) % modc
		}
		result[i] = tmp
	}
	return result
}

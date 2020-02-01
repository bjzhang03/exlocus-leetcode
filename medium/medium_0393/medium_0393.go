package medium_0393

func validUtf8(data []int) bool {
	result := true
	if len(data) > 0 {
		if data[0] < 128 {
			nextdata := append([]int{}, data[1:]...)
			result = validUtf8(nextdata)
		} else if data[0] >= 192 && data[0] < 224 {
			check, nextdata := validOne(data)
			result = check && validUtf8(nextdata)
		} else if data[0] >= 224 && data[0] < 240 {
			check, nextdata := validTwo(data)
			result = check && validUtf8(nextdata)
		} else if data[0] >= 240 && data[0] < 248 {
			check, nextdata := validThree(data)
			result = check && validUtf8(nextdata)
		} else {
			result = false
		}
	}
	return result
}

// 直接根据数字的范围来对数组进行判断
func validOne(data []int) (bool, []int) {
	result := true
	nextdata := []int{}

	if len(data) >= 2 {
		nextdata = append(nextdata, data[2:]...)
		for i := 1; i < 2; i++ {
			if !(data[i] >= 128 && data[i] < 192) {
				result = false
			}
		}
	} else {
		result = false
	}
	return result, nextdata
}

func validTwo(data []int) (bool, []int) {
	result := true
	nextdata := []int{}

	if len(data) >= 3 {
		nextdata = append(nextdata, data[3:]...)
		for i := 1; i < 3; i++ {
			if !(data[i] >= 128 && data[i] < 192) {
				result = false
			}
		}
	} else {
		result = false
	}
	return result, nextdata
}

func validThree(data []int) (bool, []int) {
	result := true
	nextdata := []int{}

	if len(data) >= 4 {
		nextdata = append(nextdata, data[4:]...)
		for i := 1; i < 4; i++ {
			if !(data[i] >= 128 && data[i] < 192) {
				result = false
			}
		}
	} else {
		result = false
	}
	return result, nextdata
}

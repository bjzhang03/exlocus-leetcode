package easy

func plusOne(digits []int) []int {
	result := []int{}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		result = append(result, tmp)
	}
	if carry > 0 {
		result = append(result, carry)
	}
	res := []int{}
	for i := len(result) - 1; i >= 0; i-- {
		res = append(res, result[i])
	}
	return res
}

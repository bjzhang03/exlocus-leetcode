package medium_0402

func removeKdigits(num string, k int) string {
	result := ""
	if k == 0 {
		result = num
	} else if k > 0 {
		// 先对数据进行预处理
		for {
			l := len(num)
			num, k = prehandler(num, k)
			if len(num) == l {
				break
			}
		}
		// 找到最小的一个数据
		for k > 0 {
			tmpnum := num
			for i := 0; i < len(num); i++ {
				save := "" + num[0:i] + num[i+1:]
				if isLarger(tmpnum, save) {
					tmpnum = save
				}
			}
			num = tmpnum
			k--
		}
		result = removePreZero(num)
	}
	if len(result) == 0 {
		result = "0"
	}
	return result
}

func isLarger(a string, b string) bool {
	result := true
	// 去掉前导0
	sta := 0
	for sta < len(a) && a[sta] == '0' {
		sta++
	}
	stb := 0
	for stb < len(b) && b[stb] == '0' {
		stb++
	}
	poa := a[sta:]
	pob := b[stb:]
	// 比较两个数字的大小
	if len(poa) < len(pob) {
		result = false
	} else if len(poa) == len(pob) {
		for i := 0; i < len(poa); i++ {
			if poa[i] < pob[i] {
				result = false
			} else if poa[i] > pob[i] {
				break
			}
		}
	}
	return result
}

// 先去掉很多0的数据
func prehandler(num string, k int) (string, int) {
	sti := 0
	for sti < len(num) && num[sti] == '0' {
		sti++
	}
	eni := sti
	for eni < len(num) && num[eni] != '0' {
		eni++
	}
	if k > (eni - sti) {
		num = num[:sti] + num[eni:]
		k = k - (eni - sti)
	}
	return num, k
}

func removePreZero(item string) string {
	sta := 0
	for sta < len(item) && item[sta] == '0' {
		sta++
	}
	return item[sta:]
}

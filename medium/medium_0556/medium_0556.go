package medium_0556

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	result := -1
	if nstr := strconv.Itoa(n); len(nstr) > 0 {
		rstring := nstr
		for i := len(nstr) - 1; i >= 0; i-- {
			if i-1 >= 0 && nstr[i-1] < nstr[i] {
				// 前半部分的数据不变动
				rstring = nstr[:i-1]
				// 对后半部分的数据进行处理
				rstring = rstring + func(b string) string {
					if len(b) > 1 {
						// 当前的数字
						current := b[0]
						tari := 1
						tarb := b[1]
						// 找到比当前的数字大的最小的数字
						for i := 1; i < len(b); i++ {
							if b[i] > current && b[i] < tarb {
								tari = i
								tarb = b[i]
							}
						}
						// 对后面的数据进行排序
						save := []int{}
						for i := 0; i < len(b); i++ {
							if i != tari {
								save = append(save, int(b[i]))
							}
						}
						sort.Ints(save)
						// 转化为string,这里转化成string没有颁发直接string了，需要通过这种方式进行处理了
						result := fmt.Sprintf("%c", b[tari])
						for i := 0; i < len(save); i++ {
							result = result + fmt.Sprintf("%c", save[i])
						}
						return result
					}
					// 只有一个数字,就直接返回了
					return b
				}(nstr[i-1:])
				break
			}
		}
		// 如果找到则返回,没有发生变化则返回-1
		if rstring != nstr {
			return func(s string) int {
				r, _ := strconv.Atoi(s)
				// 数字不能超出范围
				if r < math.MaxInt32 {
					return r
				}
				return -1
			}(rstring)
		}
	}
	return result
}

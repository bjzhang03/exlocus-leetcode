package medium_0539

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 0 {
		// 先对数据进行排序
		sort.Sort(&stime{timePoints})
		// 添加第一个第一个数字换到第二天
		timePoints = append(timePoints, func(s string) string {
			// 跨到第二天的计算方法
			hra, _ := strconv.Atoi(s[:strings.Index(s, ":")])
			return strconv.Itoa(hra+24) + s[strings.Index(s, ":"):]
		}(timePoints[0]))

		// 1441=24*60+1
		result := 1441
		for i := 1; i < len(timePoints); i++ {
			if subr := timeSubtract(timePoints[i-1], timePoints[i]); result > subr {
				result = subr
			}
		}
		return result
	}
	return 0
}

func timeSubtract(as, bs string) int {
	// as 是较小的,bs是较大的
	hra, _ := strconv.Atoi(as[:strings.Index(as, ":")])
	hrb, _ := strconv.Atoi(bs[:strings.Index(bs, ":")])
	mra, _ := strconv.Atoi(as[strings.Index(as, ":")+1:])
	mrb, _ := strconv.Atoi(bs[strings.Index(bs, ":")+1:])

	return func(ha, hb, ma, mb int) int {
		if mb >= ma {
			return mb - ma + 60*(hb-ha)
		} else {
			return 60*(hb-ha) - (ma - mb)
		}

		return 0
	}(hra, hrb, mra, mrb)
}

type stime struct {
	data []string
}

func (st stime) Len() int {
	return len(st.data)
}

func (st stime) Swap(i, j int) {
	tmp := st.data[i]
	st.data[i] = st.data[j]
	st.data[j] = tmp
}

func (st stime) Less(i, j int) bool {
	return func(as, bs string) bool {
		// 针对时的比较
		hra, _ := strconv.Atoi(as[:strings.Index(as, ":")])
		hrb, _ := strconv.Atoi(bs[:strings.Index(bs, ":")])
		if hra < hrb {
			return true
		} else if hra > hrb {
			return false
		}
		// 针对分的比较
		mra, _ := strconv.Atoi(as[strings.Index(as, ":")+1:])
		mrb, _ := strconv.Atoi(bs[strings.Index(bs, ":")+1:])
		if mra <= mrb {
			return true
		} else {
			return false
		}
	}(st.data[i], st.data[j])
}

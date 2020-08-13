package medium_0636

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	if len(logs) > 0 && n > 0 {
		save := map[int]int{}
		restack := []cpuinterval{}
		used := map[string]int{}
		for i := 0; i < len(logs); i++ {
			current := logToCpuinterval(i, logs[i])
			if current.flag == "start" {
				restack = append(restack, current)
			} else if current.flag == "end" {
				if _, ok := save[current.pronum]; ok {
					save[current.pronum] += current.tstamp - restack[len(restack)-1].tstamp + 1 - suminterval(logs[restack[len(restack)-1].index+1:current.index], &used)
				} else if !ok {
					save[current.pronum] = current.tstamp - restack[len(restack)-1].tstamp + 1 - suminterval(logs[restack[len(restack)-1].index+1:current.index], &used)
				}
				restack = restack[:len(restack)-1]
			}
		}

		result := []int{}
		for i := 0; i < n; i++ {
			result = append(result, save[i])
		}
		return result
	}
	return nil
}

// 计算内部的interval的和
func suminterval(logs []string, used *map[string]int) int {
	result := 0
	logstr := fmt.Sprint(logs)
	if _, ok := (*used)[logstr]; ok {
		return (*used)[logstr]
	} else if !ok && len(logs) > 0 {
		restack := []cpuinterval{}
		for i := 0; i < len(logs); i++ {
			current := logToCpuinterval(i, logs[i])
			if current.flag == "start" {
				restack = append(restack, current)
			} else if current.flag == "end" {
				result = result + current.tstamp - restack[len(restack)-1].tstamp + 1 - suminterval(logs[restack[len(restack)-1].index+1:current.index], used)
				restack = restack[:len(restack)-1]
			}
		}
	}
	(*used)[logstr] = result
	return result
}

// cpuinterval的数据
type cpuinterval struct {
	index  int
	pronum int
	flag   string
	tstamp int
}

// 把log的数据转化成cpuinterval的数据
func logToCpuinterval(index int, s string) cpuinterval {
	processNum, _ := strconv.Atoi(s[:strings.Index(s, ":")])
	timestamp, _ := strconv.Atoi(s[strings.LastIndex(s, ":")+1:])
	return *&cpuinterval{index, processNum, s[strings.Index(s, ":")+1 : strings.LastIndex(s, ":")], timestamp}
}

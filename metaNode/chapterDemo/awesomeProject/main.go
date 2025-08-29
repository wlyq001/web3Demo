package main

import (
	"fmt"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(intervals [][]int) [][]int {
	stackL := []int{}
	stackR := []int{}
	for _, v := range intervals {
		if len(stackL) == 0 {
			stackL = append(stackL, v[0])
			stackR = append(stackR, v[1])
			continue
		}
		small := v[0]
		big := v[1]
		for i := 0; i < len(stackL); i++ {
			if small > stackL[i] && small < stackR[i] {
				if big < stackR[i] {
					break
				}
				small = stackL[i]
				stackL[i] = -1
				stackR[i] = -1
			}
			if big > stackL[i] && big < stackR[i] {
				if small > stackL[i] {
					break
				}
				big = stackR[i]
				stackL[i] = -1
				stackR[i] = -1
			}
		}
		stackL = append(stackL, small)
		stackR = append(stackR, big)
	}

	result := [][]int{}
	for i := 0; i < len(stackL); i++ {
		if stackL[i] == -1 {
			continue
		}
		result = append(result, []int{stackL[i], stackR[i]})
	}
	return result
}

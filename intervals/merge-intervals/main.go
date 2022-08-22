// Merge intervals in golang language
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		// edge cases
		if len(intervals[i]) == 0 && len(intervals[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(intervals[i]) == 0 || len(intervals[j]) == 0 {
			return len(intervals[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}

		// both slices len() > 0, so can test this now:
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for i := range intervals {
		if len(res) > 0 {
			last := len(res) - 1
			if res[last][1] >= intervals[i][0] {
				res[last] = []int{min(res[last][0], intervals[i][0]), max(res[last][1], intervals[i][1])}
			} else {
				res = append(res, intervals[i])
			}
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	in := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(in))
	in2 := [][]int{{1, 4}, {0, 0}}
	fmt.Println(merge(in2))
}

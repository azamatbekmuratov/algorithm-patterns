// Non-overlapping Intervals in Golang language.
// Given an array of intervals intervals where intervals[i] = [starti, endi],
// return the minimum number of
// intervals you need to remove to make the rest of the intervals non-overlapping.
package main

import (
	"fmt"
	"sort"
)

func customSort(intervals [][]int) {
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
}

func eraseOverlapIntervals(intervals [][]int) int {
	customSort(intervals)
	res := 0
	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= prevEnd {
			prevEnd = intervals[i][1]
		} else {
			res += 1
			prevEnd = min(intervals[i][1], prevEnd)
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

func main() {
	in := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(in))
}

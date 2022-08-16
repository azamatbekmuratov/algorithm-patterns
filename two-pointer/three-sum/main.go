// 3Sum Three sum solution in Golang
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := v + nums[l] + nums[r]
			if threeSum > 0 {
				r -= 1
			} else if threeSum < 0 {
				l += 1
			} else {
				sumItems := []int{v, nums[l], nums[r]}
				res = append(res, sumItems)
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}

	return res
}

func main() {

}

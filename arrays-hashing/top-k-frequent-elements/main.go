// Top K Frequent Elements in Golang
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		if _, found := count[nums[i]]; found {
			count[nums[i]] = count[nums[i]] + 1
		} else {
			count[nums[i]] = 1
		}
	}
	for k, v := range count {
		freq[v] = append(freq[v], k)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, v := range freq[i] {
			res = append(res, v)
		}

		if len(res) == k {
			return res
		}
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{-1, -1}, 1))
}

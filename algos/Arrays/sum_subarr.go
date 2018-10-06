package arrays

import (
	"fmt"
)

//SubarraySum calculates the number of continuous subarrays that sum to a certain number
func SubarraySum() {
	t := []int{-1, -1, 1}
	fmt.Printf("%d == 1 \n", goodSolution(t, 0))

	t = []int{-1, -1, 1}
	fmt.Printf("%d == 1 \n", goodSolution(t, 1))

	t = []int{1, -1, 1, -1, 1, 1, 2}
	fmt.Printf("%d == 4 \n", goodSolution(t, 2))

	t = []int{1, 1, 1}
	fmt.Printf("%d == 2 \n", goodSolution(t, 2))

	t = []int{1}
	fmt.Printf("%d == 0 \n", goodSolution(t, 0))
}

//solution from https://leetcode.com/problems/subarray-sum-equals-k/discuss/102121/C++-prefix-sum-+-map
//avoid duplicated keys in case of 0 and also doubling because of jumps (no continous arrays as in my solution)
func goodSolution(nums []int, sum int) int {
	result := 0
	sumMap := make(map[int]int)
	cum := 0

	sumMap[0] = 1

	for _, v := range nums {
		cum += v
		result += sumMap[cum-sum]
		sumMap[cum]++
	}

	return result
}

//Not working . Is a bad ideea to put them in a map cause they might get duplicated for certain inputs.
func sSubarraySum(nums []int, sum int) int {
	result := 0
	sumMap := make(map[int]int)

	cum := 0

	for i, v := range nums {
		cum += v
		nums[i] = cum
		sumMap[cum-sum]++
	}

	if c, ok := sumMap[0]; ok {
		result += c
	}
	for _, v := range nums {
		if c, ok := sumMap[v]; ok {
			if sum == 0 {
				result += c / 2
				delete(sumMap, v)
			} else {
				result += c
			}
		}
	}

	return result
}

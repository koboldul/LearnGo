package arrays

import (
	"fmt"
)

//SubarraySum calculates the number of continuous subarrays that sum to a certain number
func SubarraySum() {
	t := []int{1, -1, 1, -1, 1, 1, 2}

	fmt.Println(sSubarraySum(t, 2))
}

func sSubarraySum(nums []int, sum int) int {
	result := 0
	sumMap := make(map[int]int)

	cum := 0

	for i, v := range nums {
		cum += v
		nums[i] = cum
		sumMap[sum-cum] = i
	}

	for _, v := range nums {
		if _, ok := sumMap[v]; ok {
			result++
		}
	}

	return result
}

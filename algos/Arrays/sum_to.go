package arrays

import (
	"fmt"
)

//TwoSum search in an array 2 numbers that add up to another
func TwoSum() {
	t := []int{3, 2, 4}

	fmt.Println(sTwoSum(t, 6))
}

func sTwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	result := make([]int, 2)

	for i, v := range nums {
		m[target-v] = i
	}

	for i, v := range nums {
		if val, ok := m[v]; ok && val != i {
			result[0] = val
			result[1] = i
			break
		}
	}

	return result
}

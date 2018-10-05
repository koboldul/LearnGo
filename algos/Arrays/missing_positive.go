package arrays

// FirstMissingPositive gets the first missing non zero positive number that is not found in the array. HARD
func FirstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			j := nums[i]
			for j <= len(nums) && j > 0 && nums[j-1] != j {
				temp := nums[j-1]
				nums[j-1] = j
				j = temp
			}
		}
	}

	idx := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			idx = i + 1
			break
		}
	}

	return idx
}

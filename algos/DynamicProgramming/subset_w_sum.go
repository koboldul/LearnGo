package dynamic_programming

import "fmt"

/*
* Given a set of numbers can a subset be formed in such a way that
* the sum of the subset is a certain value?
 */
func SubsetToSum() {
	var a = []int{4, 1, 12, 44, 9, 2}

	fmt.Println(get_subset(a, 24))
}

func get_subset(a []int, sum int) []int {
	l := len(a)
	//Table containing values that represent if a sum can be achieved
	//From a subset of length i-1
	var partial = make([]bool, (sum+1)*(l+1))
	var result = make([]int, 0, 10)

	for s := 0; s <= sum; s++ {
		for i := 0; i <= l; i++ {
			row := s * (l + 1)
			switch {
			case s == 0:
				partial[i] = true
			case s > 0 && i == 0:
				partial[row] = false
			case s < a[i-1]:
				partial[row+i] = partial[row+i-1]
			default:
				partial[row+i] = partial[row+i-1] || partial[(s-a[i-1])*(l+1)+i]
			}
		}
	}

	fmt.Println(partial[(l+1)*sum+l])
	return result

}

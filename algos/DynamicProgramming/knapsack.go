package dynamic_programming

import (
	"fmt"

	"../Utility"
)

/*
* Classic Knapsack with andd without repetition
 */
func Knapsack() {
	var b = []knapItem{knapItem{15, 15}, knapItem{12, 10}, knapItem{10, 8}, knapItem{5, 1}}
	fmt.Println(solve_no_repeat(b, 22))

	var a = []knapItem{knapItem{2, 8}, knapItem{4, 9}, knapItem{6, 13}}
	fmt.Println(solve_w_repeat(a, 22))
}

func solve_no_repeat(a []knapItem, sackLimit int) int {
	l := len(a)
	ts, tv := 0, 0

	for _, v := range a {
		ts += v.weight
		tv += v.value
	}

	if ts < sackLimit {
		return tv
	}

	var partial = make([]int, (sackLimit+1)*(l+1))

	for w := 0; w <= sackLimit; w++ {
		for i := 0; i <= l; i++ {
			row := w * (l + 1)

			switch {
			case w == 0:
				partial[i] = 0
			case i == 0:
				partial[w*(l+1)] = 0
			case a[i-1].weight > w:
				partial[row+i] = partial[row+i-1]
			default:
				new_val := partial[(w-a[i-1].weight)*(l+1)+i] + a[i-1].value
				prev_val := partial[row+i-1]
				partial[row+i] =
					utility.Max(prev_val, new_val)
			}

		}
	}

	return partial[sackLimit*(l+1)+l]
}

func solve_w_repeat(a []knapItem, sackLimit int) int {
	partial := make([]int, sackLimit+1)

	for w := 0; w <= sackLimit; w++ {
		for _, v := range a {
			switch {
			case w == 0:
				partial[w] = 0
			case v.weight > w:
				partial[w] = partial[w-1]
			default:
				partial[w] =
					utility.Max(partial[w-1], partial[w-v.weight]+v.value)
			}
		}
	}

	return partial[sackLimit]
}

type knapItem struct {
	weight int
	value  int
}

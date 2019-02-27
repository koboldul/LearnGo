package misc

import "fmt"

type Factor struct {
	F   int64
	Pow int64
}

func XPrime() {
	fmt.Println(getPrimes(2))
	fmt.Println(getPrimes(34))
	fmt.Println(getPrimes(1))
	fmt.Println(getPrimes(0))
	fmt.Println(getPrimes(3500))
}

func getPrimes(number int64) (result []Factor) {
	div := int64(2)
	currentFact := []int64{}

	for number > 1 {
		if number%div == 0 {
			number = number / div
			currentFact = append(currentFact, div)
		} else {
			switch {
			case div == 2:
				div++
			default:
				div += 2
			}
		}

		l := len(currentFact)
		if l > 0 && currentFact[0] != currentFact[l-1] {
			result = append(result, *composeResult(currentFact, l))
			currentFact = currentFact[l-1:]
		}
	}

	l := len(currentFact)
	if l > 0 {
		result = append(result, *composeResult(currentFact, l+1))
	}

	return
}

func composeResult(currentFact []int64, l int) (f *Factor) {
	f = &Factor{
		F:   currentFact[0],
		Pow: int64(l - 1),
	}

	return
}

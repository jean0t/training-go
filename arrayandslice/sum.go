package arrayandslice

func Sum(numbers []int) int {
	sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	//lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

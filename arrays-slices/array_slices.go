package array_slices

func Sum(numbers [5]int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumSlice(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(slices ...[]int) (sums []int) {
	sums = make([]int, 0, len(slices))
	for _, v := range slices {
		sums = append(sums, SumSlice(v))
	}
	return
}

func SumAllTails(slices ...[]int) (sums []int) {
	sums = make([]int, 0, len(slices))
	for _, v := range slices {
		if len(v) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, SumSlice(v[1:]))
	}
	return
}

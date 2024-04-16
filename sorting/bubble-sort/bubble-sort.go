package bubble_sort

func bubbleSort(numbers []int, isAscending bool) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if isAscending {
				if numbers[i] > numbers[j] {
					numbers[i], numbers[j] = numbers[j], numbers[i]
				}
			} else {
				if numbers[i] < numbers[j] {
					numbers[i], numbers[j] = numbers[j], numbers[i]
				}
			}
		}
	}
}

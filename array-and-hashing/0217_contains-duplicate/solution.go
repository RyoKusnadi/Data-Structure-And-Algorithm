package contains_duplicate

func containsDuplicate(nums []int) bool {
	// Condition
	// Valid: If Any Value Appears At Least Twice
	visitedArray := make(map[int]bool, 0)
	for _, num := range nums {
		if _, isExist := visitedArray[num]; isExist {
			return true
		}
		visitedArray[num] = true
	}
	return false
}

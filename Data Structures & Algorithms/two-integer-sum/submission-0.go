func twoSum(nums []int, target int) []int {
    checker := make(map[int]int)

	for index, num := range nums {
		low := target - num

		if v, ok := checker[low]; ok {
			return []int{v, index}
		}

		checker[num] = index
	}

	return []int{}
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num]++
	}

	freqMap := make(map[int][]int)

	for num, amount := range numMap {
		freqMap[amount] = append(freqMap[amount], num)
	}

	result := make([]int, 0, k)

	for i := len(nums)+1; i>0; i-- {
		for _, num := range freqMap[i] {

			result = append(result, num)

			if len(result) == k {
				return result
			}

		}

	}
	return result
}

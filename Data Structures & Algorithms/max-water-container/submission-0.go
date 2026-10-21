func maxArea(heights []int) int {
	left := 0
	right := len(heights)-1

	result := 0

	for left < right {
		width := right - left

		amount := min(heights[left], heights[right]) * width

		if amount > result {
			result = amount
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

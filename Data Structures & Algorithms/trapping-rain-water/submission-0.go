func trap(height []int) int {
	result := 0
	left := 0
	right := len(height)-1

	leftMax := height[left]
	rightMax := height[right]

	for left < right {
		if leftMax < rightMax {
			left++
			// считаем и двигаем левую границу
			if leftMax - height[left] <= 0 {
				leftMax = height[left]
			}

			result += leftMax - height[left]

		} else {
			right--
			// считаем и двигаем правую границу
			if rightMax - height[right] <= 0 {
				rightMax = height[right]
			}
			result += rightMax - height[right]
		}
	}

	return result
}

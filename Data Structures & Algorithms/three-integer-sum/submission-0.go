func threeSum(nums []int) [][]int {
	result := make([][]int, 0, 3)
	sort.Ints(nums)

	for i:=0; i<=len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}

		if i!=0 && nums[i] == nums[i-1] {
			continue
		}

		l := i+1
		r := len(nums)-1

		for l<r {
			numsSum := nums[i]+nums[l]+nums[r]

			if numsSum < 0 {
				l++
			}

			if numsSum > 0 {
				r--
			}

			if numsSum==0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				for l<r && nums[l] == nums[l-1] {
					l++
				}

				for l<r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return result
}

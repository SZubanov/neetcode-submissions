func hasDuplicate(nums []int) bool {
    registry := make(map[int]struct{})

    for _, num := range nums {
        if _, ok := registry[num]; ok {
            return true
        }

        registry[num] = struct{}{}
    }

    return false
}

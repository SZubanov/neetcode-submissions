func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)

	for _, str := range strs {
		var counter [26]int

		for _, letter := range str {
			counter[letter - 'a']++
		}
		anagramMap[counter] = append(anagramMap[counter], str)
	}

	var result [][]string

	for	_, anagram := range anagramMap {
		result = append(result, anagram)
	}

	return result
}

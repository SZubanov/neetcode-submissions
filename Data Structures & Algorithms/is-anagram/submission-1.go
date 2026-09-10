func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for i := 0; i<len(s); i++ {
		mapS[rune(s[i])]++
		mapT[rune(t[i])]++
	}

	for letter, amount := range mapS {
		if mapT[letter] != amount {
			return false
		}
	}

	return true
}

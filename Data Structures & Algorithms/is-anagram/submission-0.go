func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	mapSs := make(map[byte]int)
	mapT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mapSs[s[i]]++
		mapT[t[i]]++
	}

	for letter, amount := range mapSs {
		if mapT[letter] != amount {
			return false
		}
	}

	return true
}

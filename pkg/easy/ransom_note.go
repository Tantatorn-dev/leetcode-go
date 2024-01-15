package easy

// Problem: https://leetcode.com/problems/ransom-note
func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMap := make(map[string]int)
	magazineMap := make(map[string]int)

	for _, v := range ransomNote {
		vStr := string(v)

		if _, ok := ransomNoteMap[vStr]; ok {
			ransomNoteMap[vStr]++
			continue
		}

		ransomNoteMap[vStr] = 1
	}

	for _, v := range magazine {
		vStr := string(v)

		if _, ok := magazineMap[vStr]; ok {
			magazineMap[vStr]++
			continue
		}

		magazineMap[vStr] = 1
	}

	ret := true

	for key, val := range ransomNoteMap {
		if val > magazineMap[key] {
			return false
		}
	}

	return ret
}

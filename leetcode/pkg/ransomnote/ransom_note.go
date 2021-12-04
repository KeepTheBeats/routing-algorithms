package ransomnote

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var ransomNoteMap map[byte]int = make(map[byte]int)
	var magazineMap map[byte]int = make(map[byte]int)

	for _, letter := range []byte(ransomNote) {
		ransomNoteMap[letter]++
	}

	for _, letter := range []byte(magazine) {
		magazineMap[letter]++
	}

	for letter, numberRansomNote := range ransomNoteMap {
		if numberRansomNote > magazineMap[letter] {
			return false
		}
	}
	return true
}

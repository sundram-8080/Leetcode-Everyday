package problems

func ValidAnangram(source string, target string) bool {

	charCount := make(map[rune]int)

	for _, value := range source {
		charCount[value]++
	}

	for _, value := range target {
		charCount[value]--
	}

	for _, value := range charCount {
		if value != 0 {
			return false
		}
	}

	return true

}

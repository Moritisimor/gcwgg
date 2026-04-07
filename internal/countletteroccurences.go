package internal

func countLetterOccurrences(word string) map[string]int {
	occurrences := map[string]int{}
	for _, c := range word {
		letter := string(c)
		occurrences[letter]++
	}

	return occurrences
}

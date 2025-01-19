package processor

// MergeResults объединяет результаты из нескольких карт.
func MergeResults(results <-chan map[string]int) map[string]int {
	finalFreq := make(map[string]int)

	for result := range results {
		for word, count := range result {
			finalFreq[word] += count
		}
	}

	return finalFreq
}

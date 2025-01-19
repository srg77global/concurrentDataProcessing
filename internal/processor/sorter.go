package processor

import (
	"sort"
)

type WordCount struct {
	Word  string
	Count int
}

// GetTopWords возвращает топ 10 слов.
func GetTopWords(wordFreq map[string]int) []WordCount {
	wordCounts := make([]WordCount, 0, len(wordFreq))

	for word, count := range wordFreq {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	if len(wordCounts) > 10 {
		return wordCounts[:10]
	}
	return wordCounts
}

package processor

import (
	"strings"
	"unicode"
)

// ProcessChunk подсчитывает слова в чанке.
func ProcessChunk(chunk string, results chan<- map[string]int) {
	wordFreq := make(map[string]int)
	words := strings.FieldsFunc(chunk, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		word = strings.ToLower(word)
		wordFreq[word]++
	}

	results <- wordFreq
}

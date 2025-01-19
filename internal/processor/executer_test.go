package processor

import (
	"os"
	"sort"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Создаем временный файл.
	content := "hello world hello go go go\nparallel programming in go\n"
	tempFile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Записываем данные в файл.
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Тестируем ProcessFile.
	topWords, err := ProcessFile(tempFile.Name(), 5)
	if err != nil {
		t.Fatalf("failed to process file: %v", err)
	}

	// Сортируем результат по убыванию частоты.
	sort.Slice(topWords, func(i, j int) bool {
		return topWords[i].Count > topWords[j].Count
	})

	// Проверяем результат.
	expected := []WordCount{
		{"go", 4},
		{"hello", 2},
		{"world", 1},
		{"parallel", 1},
		{"programming", 1},
		{"in", 1},
	}

	for i, word := range expected {
		if i >= len(topWords) {
			t.Errorf("expected %v, but got fewer words", word)
			continue
		}
		if word.Word != topWords[i].Word || word.Count != topWords[i].Count {
			t.Errorf("expected %v, got %v", word, topWords[i])
		}
	}
}

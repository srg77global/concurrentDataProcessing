package processor

import (
	"os"
	"testing"
)

func TestReadChunksWithProgress(t *testing.T) {
	// Создаем временный файл.
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed creating temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Записываем данные в файл.
	content := "hello world this is a test of the scanner functionality"
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("failed write data to temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	// Создаем канал для результатов.
	chunks := make(chan string)

	// Запускаем функцию в отдельной горутине.
	go func() {
		err := ReadChunksWithProgress(tempFile.Name(), chunks)
		if err != nil {
			t.Errorf("error execution ReadChunksWithProgress: %v", err)
		}
	}()

	// Считываем результаты из канала.
	var result []string
	for chunk := range chunks {
		result = append(result, chunk)
	}

	// Проверяем результат.
	expected := []string{"hello", "world", "this", "is", "a", "test", "of", "the", "scanner", "functionality"}

	if len(result) != len(expected) {
		t.Errorf("expected %d words, got %d", len(expected), len(result))
	}

	for i, word := range expected {
		if result[i] != word {
			t.Errorf("expected word %q, got %q position %d", word, result[i], i)
		}
	}
}

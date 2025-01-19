package processor

import (
	"reflect"
	"testing"
)

func TestProcessChunk(t *testing.T) {
	// Создаем чанк и канал.
	chunk := "hello world hello Insit"
	results := make(chan map[string]int, 1)

	// Тестируем обработку чанка.
	ProcessChunk(chunk, results)
	close(results)

	// Проверяем результат.
	expected := map[string]int{"hello": 2, "world": 1, "insit": 1}
	for res := range results {
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("expected %v, got %v", expected, res)
		}
	}
}

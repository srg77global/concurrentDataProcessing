package processor

import (
	"reflect"
	"testing"
)

func TestMergeResults(t *testing.T) {
	// Имитируем запись в канал.
	results := make(chan map[string]int, 3)
	results <- map[string]int{"hello": 2, "insit": 1}
	results <- map[string]int{"go": 3, "hello": 1}
	results <- map[string]int{"test": 4}
	close(results)

	// Собираем данные из канала в мапу.
	actual := MergeResults(results)

	// Проверяем результат.
	expected := map[string]int{"hello": 3, "insit": 1, "go": 3, "test": 4}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

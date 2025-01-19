package processor

import (
	"reflect"
	"testing"
)

func TestGetTopWords(t *testing.T) {
	// Создаем мапу со значениями для сортировки.
	wordFreq := map[string]int{
		"hello": 5,
		"world": 3,
		"go":    10,
		"test":  7,
		"code":  2,
		"run":   4,
	}

	// Сортируем мапу.
	actual := GetTopWords(wordFreq)

	// Проверяем результат.
	expected := []WordCount{
		{"go", 10},
		{"test", 7},
		{"hello", 5},
		{"run", 4},
		{"world", 3},
		{"code", 2},
	}
	if !reflect.DeepEqual(actual, expected[:len(actual)]) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

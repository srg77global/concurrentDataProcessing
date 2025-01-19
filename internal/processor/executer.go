package processor

import "sync"

// ProcessFile - основная функция, связывает модули и управляет потоком данных.
func ProcessFile(filePath string, workerCount int) ([]WordCount, error) {
	// Создаем каналы для передачи данных.
	chunks := make(chan string, workerCount)
	results := make(chan map[string]int, workerCount)

	// Считыванием файл чанками с прогрессом в отдельной горутине.
	var readErr error
	go func() {
		readErr = ReadChunksWithProgress(filePath, chunks)
	}()

	// Запускем обработчиков чанков (горутины).
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for chunk := range chunks {
				ProcessChunk(chunk, results)
			}
		}()
	}

	// Закрываем канал результатов после завершения обработки.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Собираем результаты.
	wordFreq := MergeResults(results)

	// Проверяем на ошибки при чтении файла.
	if readErr != nil {
		return nil, readErr
	}

	// Сортируем и получаем топ 10 слов.
	return GetTopWords(wordFreq), nil
}

package processor

import (
	"bufio"
	"fmt"
	"os"
)

// ReadChunksWithProgress читает файл и отправляет слова в канал, отображая прогресс.
func ReadChunksWithProgress(filePath string, chunks chan<- string) error {
	defer close(chunks)

	// Открываем файл.
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Получаем общий размер файла.
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	totalSize := fileInfo.Size()

	// Инициализируем сканнер и устанавливаем разделитель для чтения по словам.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var bytesRead int64

	// Сканируем текст и отправляем в канал.
	for scanner.Scan() {
		word := scanner.Text()
		chunks <- word

		// Обновляем прогресс
		bytesRead += int64(len(word) + 1) // +1 для пробела.
		printProgress(bytesRead, totalSize)
	}

	// Проверяем ошибки сканирования.
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("\nReading succeed.")
	return nil
}

// printProgress выводит прогресс в процентах.
func printProgress(current, total int64) {
	percentage := float64(current) / float64(total) * 100
	fmt.Printf("\rProgress: %.2f%%", percentage)
}

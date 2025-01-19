package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/srg77global/concurrentDataProcessing/internal/logger"
	"github.com/srg77global/concurrentDataProcessing/internal/processor"
)

func main() {
	// Параметры запуска.
	filePath := flag.String("file", "", "path to the file")
	workerCount := flag.Int("workers", 4, "number of goroutines")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("empty path to the file; use flag -file.")
	}

	// Инициализация логгера.
	logger.InitLogger()

	// Запуск обработки файла.
	topWords, err := processor.ProcessFile(*filePath, *workerCount)
	if err != nil {
		log.Fatalf("file process error: %v", err)
	}

	// Вывод результата.
	fmt.Println("TOP 10 words by frequency:")
	for i, word := range topWords {
		fmt.Printf("%d. %s: %d\n", i+1, word.Word, word.Count)
	}
}

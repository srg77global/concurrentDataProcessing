Описание проекта.
Эта программа читает текстовый файл большого размера, подсчитывает частоту слов с использованием параллельной обработки, и выводит топ 10 самых часто встречающихся слов.

Инструкция по запуску.
go run cmd/main.go -file=testdata/testfile.txt -workers=4

Примеры использования.
Запуск с различными параметрами:
go run cmd/main.go -file=file.txt -workers=5
go run cmd/main.go -file=file.txt
go run cmd/main.go -file=path/to/the/file.txt -workers=10

Архитектура.
- Чтение чанков файла — reader.go
- Параллельная обработка — worker.go
- Объединение результатов — counter.go
- Сортировка слов — sorter.go
- Объединение и выполнение функций модулей  - executer.go

Используемые технологии.
- Go standard library.
- Горутины и каналы.# concurrentDataProcessing

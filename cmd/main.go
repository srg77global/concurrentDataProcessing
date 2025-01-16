package main

import (
	"concurrentDataProcessing/pkg/opener"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../files/metro.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	dictionary := make(map[string]int)

	opener.ReadLines(file, dictionary)

	for key, value := range dictionary {
		fmt.Printf("%s: %d\n", key, value)
	}
	print("Привет мир!")
}

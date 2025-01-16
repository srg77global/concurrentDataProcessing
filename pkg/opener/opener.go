package opener

import (
	"bufio"
	"os"
)

func ReadLines(desc *os.File, dictionary map[string]int) (map[string]int, error) { //
	scanner := bufio.NewScanner(desc)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if _, ok := dictionary[scanner.Text()]; ok {
			dictionary[scanner.Text()] += 1
		} else {
			dictionary[scanner.Text()] = 1
		}
	}
	return dictionary, nil
}

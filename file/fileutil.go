package file

import (
	"bufio"
	"os"
)

func readLines(input string) ([]string, error) {
	inFile, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	slice := []string{}

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return slice, nil
}

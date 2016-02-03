package file

import (
	"log"
)

func readLines(input string) ([]string, error) {
	inFile, err := os.Open(urlFile)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	slice := []string{}

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// log.Fatal(err)
		fmt.Println("scanner Err : ", err)
		return nil, err
	}

	return slice, nil
}

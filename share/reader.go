package share

import (
	"bufio"
	"os"
)

func ReadByLine(fileName string, lineHandler func(line string)) {
	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	// prepare the scanner to scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// get the line from the file
		line := scanner.Text()

		lineHandler(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

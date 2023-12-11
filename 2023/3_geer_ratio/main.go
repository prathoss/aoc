package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	partNumberSum := SumPartNumbers("input.txt")
	log.Printf("Sum of part numbers: %d\n", partNumberSum)
}

func SumPartNumbers(fileName string) int {
	usedPartNumbers := map[Coordinates]int{}
	ReadByLine(fileName, func(previousLine, currentLine, nextLine string, currentLineNumber int) {
		for i, r := range currentLine {
			if !IsSymbol(r) {
				continue
			}

			storeToUsedPartNumbers := func(number int, startIndex int) {
				c := Coordinates{
					Row:    currentLineNumber,
					Column: startIndex,
				}
				if _, alreadyUsed := usedPartNumbers[c]; !alreadyUsed {
					usedPartNumbers[c] = number
				}
			}

			handleRune := func(line []rune, index int) {
				if IsNumeric(line[index]) {
					number, startIndex := GetNumberWithStartIndex(line, index)
					storeToUsedPartNumbers(number, startIndex)
				}
			}

			handleLine := func(line string) {
				runeLine := []rune(line)
				if line != "" {
					if i > 0 {
						handleRune(runeLine, i-1)
					}

					handleRune(runeLine, i)

					if i < len(line)-1 {
						handleRune(runeLine, i+1)
					}
				}
			}

			handleLine(previousLine)
			handleLine(currentLine)
			handleLine(nextLine)
		}
	})

	partNumberSum := 0
	for _, partNumber := range usedPartNumbers {
		partNumberSum += partNumber
	}
	return partNumberSum
}

type Coordinates struct {
	Row    int
	Column int
}

func ReadByLine(fileName string, lineHandler func(previousLine, currentLine, nextLine string, currentLineNumber int)) {
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

	lineNumber := -1
	var previous, current, next string

	shiftLines := func(newNext string) {
		previous = current
		current = next
		next = newNext

		if lineNumber >= 0 {
			lineHandler(previous, current, next, lineNumber)
		}
	}
	// prepare the scanner to scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		shiftLines(scanner.Text())

		lineNumber += 1
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	shiftLines("")
}

func IsSymbol(r rune) bool {
	if r == '.' {
		return false
	}

	if IsNumeric(r) {
		return false
	}

	return true
}

func IsNumeric(r rune) bool {
	return unicode.IsDigit(r)
}

func GetNumberWithStartIndex(line []rune, index int) (int, int) {
	startIndex := index
	endIndex := index

	for startIndex > 0 && IsNumeric(line[startIndex-1]) {
		startIndex -= 1
	}
	for endIndex < len(line)-1 && IsNumeric(line[endIndex+1]) {
		endIndex += 1
	}

	number, err := strconv.Atoi(string(line[startIndex : endIndex+1]))
	if err != nil {
		panic(err)
	}

	return number, startIndex
}

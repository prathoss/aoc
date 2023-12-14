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
	sum := 0
	ReadByLine(fileName, func(previousLine, currentLine, nextLine string, currentLineNumber int) {
		i := 0
		for i < len(currentLine) {
			r := rune(currentLine[i])
			if !IsNumeric(r) {
				i += 1
				continue
			}

			num, startIndex, endIndex := GetNumberWithStartAndEndIndex([]rune(currentLine), &i)
			contains := false
			for _, searchLine := range []string{previousLine, currentLine, nextLine} {
				if searchLine == "" {
					continue
				}
				if ContainSymbol(searchLine[max(0, startIndex-1):min(len(searchLine), endIndex+1)]) {
					contains = true
					break
				}
			}
			if contains {
				sum += num
			}
			i += 1
		}
	})

	return sum
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

func GetNumberWithStartAndEndIndex(line []rune, i *int) (int, int, int) {
	startIndex := *i
	for *i < len(line) && IsNumeric(line[*i]) {
		*i += 1
	}

	number, err := strconv.Atoi(string(line[startIndex:*i]))
	if err != nil {
		panic(err)
	}
	return number, startIndex, *i
}

func ContainSymbol(s string) bool {
	for _, r := range s {
		if IsSymbol(r) {
			return true
		}
	}
	return false
}

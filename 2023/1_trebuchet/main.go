package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		slog.Error("could not open file", "err", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first := 10 * getFirstNumber(line)
		last := getLastNumber(line)
		sum += first + last
	}

	if err := scanner.Err(); err != nil {
		slog.Error("scanner error", "err", err)
	}

	slog.Info(strconv.Itoa(sum))
}

func getFirstNumber(s string) int {
	for i := 0; i < len(s); i++ {
		r := s[i]
		if num, err := strconv.Atoi(string(r)); err == nil {
			return num
		}
		if contains, num := getNumberFromText(s[0 : i+1]); contains {
			return num
		}
	}

	panic("could not find number")
}

func getLastNumber(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		if num, err := strconv.Atoi(string(r)); err == nil {
			return num
		}
		if contains, num := getNumberFromText(s[i:len(s)]); contains {
			return num
		}
	}

	panic("could not find number")
}

func getNumberFromText(s string) (bool, int) {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbersToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, num := range numbers {
		if strings.Contains(s, num) {
			return true, numbersToInt[num]
		}
	}
	return false, -1
}

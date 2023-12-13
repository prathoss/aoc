package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/prathoss/advent_of_code/share"
)

func main() {
	log.Printf("Scratchcard matches: %v\n", CountScratchcards("input.txt"))
	log.Printf("Scratchcard copies: %v\n", CountScratchcardsCopies("input.txt"))
}

func CountScratchcardsCopies(fileName string) int {
	copies := 0
	var cards []Card
	share.ReadByLine(fileName, func(line string) {
		cards = append(cards, NewCard(line))
	})

	for i, card := range cards {
		points := card.CountMatches()
		for j := 0; j < min(points, len(cards)-i); j++ {
			cards[i+j+1].Copies += card.Copies
		}

		copies += card.Copies
	}
	return copies
}

func CountScratchcards(fileName string) int {
	points := 0
	share.ReadByLine(fileName, func(line string) {
		card := NewCard(line)
		cardPoints := card.CountPoints()
		points += cardPoints
	})
	return points
}

type Card struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
	Copies         int
}

func (c Card) CountPoints() int {
	matches := c.CountMatches()
	return CountPointsByMatches(matches)
}

func (c Card) CountMatches() int {
	matches := 0

	for _, number := range c.Numbers {
		if slices.Contains(c.WinningNumbers, number) {
			matches += 1
		}
	}

	return matches
}

func CountPointsByMatches(matches int) int {
	if matches == 0 {
		return 0
	}

	return 1 << (matches - 1)
}

func NewCard(line string) Card {
	card := Card{
		Copies: 1,
	}
	strCard := strings.Split(line, ":")

	heading := strings.Split(strCard[0], " ")
	strId := heading[len(heading)-1]

	var err error
	card.Id, err = strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	nums := strings.Split(strCard[1], "|")
	card.WinningNumbers = parseNumbers(nums[0])
	card.Numbers = parseNumbers(nums[1])
	return card
}

func parseNumbers(s string) []int {
	result := make([]int, 0, len(s)/2)

	addNumber := func(startIndex, endIndex int) {
		number, err := strconv.Atoi(s[startIndex:endIndex])
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	startedReadingNum := false
	numberStartIndex := 0
	for i, r := range s {
		if IsNumeric(r) && !startedReadingNum {
			startedReadingNum = true
			numberStartIndex = i
		}
		if !IsNumeric(r) && startedReadingNum {
			addNumber(numberStartIndex, i)
			startedReadingNum = false
		}
	}

	if startedReadingNum {
		addNumber(numberStartIndex, len(s))
	}
	return result
}

func IsNumeric(r rune) bool {
	return unicode.IsDigit(r)
}

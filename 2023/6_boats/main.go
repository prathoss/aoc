package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/prathoss/advent_of_code/share"
)

func main() {
	product := GetMarginOfError()
	log.Printf("Number of ways to beat the record: %v", product)
	log.Printf("Margin of error: %v", GetMarginOfError2())
}

func GetMarginOfError() int {
	lines := make([]string, 0)
	share.ReadByLine("input.txt", func(line string) {
		lines = append(lines, line)
	})

	f := strings.Fields(lines[0])
	stringTimes := f[1:]
	times := make([]int, 0, len(stringTimes))
	for _, strTime := range stringTimes {
		time, err := strconv.Atoi(strTime)
		if err != nil {
			panic(err)
		}
		times = append(times, time)
	}

	f = strings.Fields(lines[1])
	stringDistances := f[1:]
	distances := make([]int, 0, len(stringDistances))
	for _, strDistance := range stringDistances {
		distance, err := strconv.Atoi(strDistance)
		if err != nil {
			panic(err)
		}
		distances = append(distances, distance)
	}

	product := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		record := distances[i]

		possibilities := 0
		for j := 1; j <= time; j++ {
			ct := j * (time - j)
			if ct > record {
				possibilities++
			}
		}
		product *= possibilities
	}
	return product
}

func GetMarginOfError2() int {
	lines := make([]string, 0)
	share.ReadByLine("input.txt", func(line string) {
		lines = append(lines, line)
	})

	f := strings.Fields(lines[0])
	strTime := strings.Join(f[1:], "")
	time, err := strconv.Atoi(strTime)
	if err != nil {
		panic(err)
	}

	f = strings.Fields(lines[1])
	strDistance := strings.Join(f[1:], "")
	distance, err := strconv.Atoi(strDistance)
	if err != nil {
		panic(err)
	}

	possibilities := 0
	for j := 1; j <= time; j++ {
		ct := j * (time - j)
		if ct > distance {
			possibilities++
		}
	}

	return possibilities
}

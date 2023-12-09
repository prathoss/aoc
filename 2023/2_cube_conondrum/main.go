package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/prathoss/advent_of_code/share"
)

var constraints = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	sum := 0
	share.ReadByLine("input.txt", func(line string) {
		game := NewGame(line)
		if game.IsPlayable(constraints) {
			sum += game.Id
		}
	})
	log.Printf("Sum of IDs: %d\n", sum)

	sumPower := 0
	share.ReadByLine("input.txt", func(line string) {
		game := NewGame(line)
		sumPower += game.GetPower()
	})
	log.Printf("Sum of power: %d\n", sumPower)
}

type Game struct {
	Id    int
	Plays []map[string]int
}

func (g Game) IsPlayable(constraints map[string]int) bool {
	for _, play := range g.Plays {
		for color, number := range play {
			maxColor := constraints[color]
			if number > maxColor {
				return false
			}
		}
	}
	return true
}

func (g Game) GetPower() int {
	required := map[string]int{}
	for _, play := range g.Plays {
		for color, num := range play {
			if r, exists := required[color]; exists {
				required[color] = max(r, num)
			} else {
				required[color] = num
			}
		}
	}
	power := 1
	for _, r := range required {
		power *= r
	}
	return power
}

func NewGame(line string) Game {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	header := strings.Split(line, ":")

	id := ExtractId(header)

	strPlays := strings.Split(header[1], ";")
	plays := make([]map[string]int, 0, len(strPlays))
	for _, strPlay := range strPlays {
		played := map[string]int{}
		strCubes := strings.Split(strPlay, ",")
		for _, strCube := range strCubes {
			number, color := ExtractCube(strCube)
			played[color] = number
		}
		plays = append(plays, played)
	}

	return Game{
		Id:    id,
		Plays: plays,
	}
}

func ExtractId(header []string) int {
	strId := strings.Split(header[0], " ")[1]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	return id
}

func ExtractCube(s string) (int, string) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	strNumber := split[0]

	number, err := strconv.Atoi(strNumber)
	if err != nil {
		panic(err)
	}

	color := split[1]

	return number, color
}

package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/prathoss/advent_of_code/share"
)

func main() {
	m := NewAlmanac("input.txt")
	minimalLocation := m.GetMinimalLocation()
	log.Printf("Minimal location: %v", minimalLocation)

	minimalLocation2 := m.GetMinimalLocation2()
	log.Printf("Minimal location2: %v", minimalLocation2)
}

type Almanac struct {
	Seeds                 []int
	seedToSoil            SrcDestMap
	soilToFertilizer      SrcDestMap
	fertilizerToWater     SrcDestMap
	waterToLight          SrcDestMap
	lightToTemperature    SrcDestMap
	temperatureToHumidity SrcDestMap
	humidityToLocation    SrcDestMap
}

func NewAlmanac(fileName string) Almanac {
	s := make([]string, 0, 261)
	share.ReadByLine(fileName, func(line string) {
		s = append(s, line)
	})

	m := Almanac{}
	actions := []func(s []string){
		func(s []string) {
			m.Seeds = ParseRequiredSeeds(s)
		},
		func(s []string) {
			m.seedToSoil = ParseMap(s)
		},
		func(s []string) {
			m.soilToFertilizer = ParseMap(s)
		},
		func(s []string) {
			m.fertilizerToWater = ParseMap(s)
		},
		func(s []string) {
			m.waterToLight = ParseMap(s)
		},
		func(s []string) {
			m.lightToTemperature = ParseMap(s)
		},
		func(s []string) {
			m.temperatureToHumidity = ParseMap(s)
		},
		func(s []string) {
			m.humidityToLocation = ParseMap(s)
		},
	}
	actionIndex := 0
	startIndex := 0
	for i, line := range s {
		if line == "" {
			actions[actionIndex](s[startIndex:i])
			actionIndex++
			startIndex = i + 1
		}
	}
	// file does not end with empty line
	actions[actionIndex](s[startIndex:])
	return m
}

func ParseRequiredSeeds(s []string) []int {
	line := s[0]
	split := strings.Split(line, ": ")
	stringSeeds := strings.Split(split[1], " ")
	seeds := make([]int, 0, len(stringSeeds))
	for _, stringSeed := range stringSeeds {
		seed, err := strconv.Atoi(stringSeed)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func ParseMap(s []string) SrcDestMap {
	s = s[1:]
	m := NewSrcDestMap()
	for _, line := range s {
		sp := strings.Split(line, " ")
		destination, err := strconv.Atoi(sp[0])
		if err != nil {
			panic(err)
		}
		source, err := strconv.Atoi(sp[1])
		if err != nil {
			panic(err)
		}
		rng, err := strconv.Atoi(sp[2])
		if err != nil {
			panic(err)
		}

		itm := NewSrcDestItem(source, destination, rng)
		m.Insert(itm)
	}
	return m
}

func (a Almanac) GetMinimalLocation() int {
	initialLocationSet := false
	minLocation := 0
	for _, seed := range a.Seeds {
		soil := a.seedToSoil.Get(seed)
		fertilizer := a.soilToFertilizer.Get(soil)
		water := a.fertilizerToWater.Get(fertilizer)
		light := a.waterToLight.Get(water)
		temp := a.lightToTemperature.Get(light)
		humidity := a.temperatureToHumidity.Get(temp)
		location := a.humidityToLocation.Get(humidity)

		if initialLocationSet {
			minLocation = min(minLocation, location)
		} else {
			minLocation = location
			initialLocationSet = true
		}
	}
	return minLocation
}

func (a Almanac) GetMinimalLocation2() int {
	requiredSeeds := a.GetRequiredSeeds()

	initialLocationSet := false
	minLocation := 0
	for _, seed := range requiredSeeds {
		soil := a.seedToSoil.Get(seed)
		fertilizer := a.soilToFertilizer.Get(soil)
		water := a.fertilizerToWater.Get(fertilizer)
		light := a.waterToLight.Get(water)
		temp := a.lightToTemperature.Get(light)
		humidity := a.temperatureToHumidity.Get(temp)
		location := a.humidityToLocation.Get(humidity)

		if initialLocationSet {
			minLocation = min(minLocation, location)
		} else {
			minLocation = location
			initialLocationSet = true
		}
	}
	return minLocation
}

func (a Almanac) GetRequiredSeeds() []int {
	requiredSeeds := make([]int, 0, len(a.Seeds))

	for i := 0; i < len(a.Seeds); i += 2 {
		start := a.Seeds[i]
		rng := a.Seeds[i+1]
		for j := 0; j < rng; j++ {
			requiredSeeds = append(requiredSeeds, start+j)
		}
	}

	return requiredSeeds
}

type SrcDestItem struct {
	Src   int
	Dest  int
	Range int
}

func NewSrcDestItem(src, dest, rng int) SrcDestItem {
	return SrcDestItem{
		Src:   src,
		Dest:  dest,
		Range: rng,
	}
}

type SrcDestMap struct {
	items []SrcDestItem
}

func NewSrcDestMap() SrcDestMap {
	return SrcDestMap{items: make([]SrcDestItem, 0)}
}

func (m *SrcDestMap) Insert(sdi SrcDestItem) {
	if len(m.items) == 0 {
		m.items = []SrcDestItem{sdi}
		return
	}
	for i := 0; i < len(m.items); i++ {
		item := m.items[i]

		if sdi.Src < item.Src {
			m.items = append(m.items[:i], append([]SrcDestItem{sdi}, m.items[i:]...)...)
			return
		}
	}

	m.items = append(m.items, sdi)
}

func (m *SrcDestMap) Get(i int) int {
	for _, item := range m.items {
		if i >= item.Src && i < item.Src+item.Range {
			increment := i - item.Src
			return item.Dest + increment
		}
	}
	return i
}

package main

import (
	"slices"
	"testing"
)

func TestGetMinLocation(t *testing.T) {

}

func TestNewMap(t *testing.T) {
	tc := []struct {
		name     string
		fileName string
		expected Almanac
	}{
		{
			name:     "sample",
			fileName: "test_sample.txt",
			expected: Almanac{
				Seeds:                 []int{79, 14, 55, 13},
				seedToSoil:            NewSrcDestMap(),
				soilToFertilizer:      NewSrcDestMap(),
				fertilizerToWater:     NewSrcDestMap(),
				waterToLight:          NewSrcDestMap(),
				lightToTemperature:    NewSrcDestMap(),
				temperatureToHumidity: NewSrcDestMap(),
				humidityToLocation:    NewSrcDestMap(),
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			m := NewAlmanac(tt.fileName)
			if !slices.Equal(m.Seeds, tt.expected.Seeds) {
				t.Fatalf("expected Seeds to be %v, but got %v", tt.expected.Seeds, m.Seeds)
			}
		})
	}
}

func TestMap_GetMinimalLocation_Sample(t *testing.T) {
	m := NewAlmanac("test_sample.txt")
	l := m.GetMinimalLocation()
	expectedLocation := 35

	if l != expectedLocation {
		t.Fatalf("GetMinimalLocation() expected to return %v, but got %v", expectedLocation, l)
	}
}

func TestMap_GetMinimalLocation2_Sample(t *testing.T) {
	m := NewAlmanac("test_sample.txt")
	l := m.GetMinimalLocation2()
	expectedLocation := 46

	if l != expectedLocation {
		t.Fatalf("GetMinimalLocation() expected to return %v, but got %v", expectedLocation, l)
	}
}

package fertilizer

import (
	"strconv"
	"strings"
)

type ReverterMap struct {
	destRangeStart  int
	rangeLength     int
	reversionAmount int
}

type Reverter struct {
	maps *[]ReverterMap
}

func Get(lines *[]string) *Reverter {
	m := []ReverterMap{}
	conv := Reverter{
		maps: &m,
	}

	for _, line := range *lines {
		if line == "" {
			continue
		}
		mapData := strings.Split(strings.TrimSpace(line), " ")
		destRangeStart, err := strconv.Atoi(mapData[0])
		if err != nil {
			panic(err)
		}
		sourceRangeStart, err := strconv.Atoi(mapData[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(mapData[2])
		if err != nil {
			panic(err)
		}
		*conv.maps = append(*conv.maps, ReverterMap{
			destRangeStart:  destRangeStart,
			rangeLength:     rangeLength,
			reversionAmount: (sourceRangeStart - destRangeStart),
		})
	}
	return &conv
}

func (r *Reverter) Revert(num int) int {
	for _, converterMap := range *r.maps {
		relativeIndex := (num - converterMap.destRangeStart)
		if relativeIndex >= 0 && relativeIndex < converterMap.rangeLength {
			return num + converterMap.reversionAmount
		}
	}
	return num
}

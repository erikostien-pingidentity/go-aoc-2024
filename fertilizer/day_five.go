package fertilizer

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/erikostien-pingidentity/go-aoc-2024/common"
)

func SolveDayFive(dataFilename string) (int, int) {
	lines := common.GetLinesFromTestdataFile(dataFilename)

	seeds, reverters := parseLineData(&lines)

	answerPartOne := solvePartOne(seeds, reverters)

	answerPartTwo := solvePartTwo(seeds, reverters)

	return answerPartOne, answerPartTwo
}

func solvePartOne(seeds *[]int, reverters *[]Reverter) int {
	for minLocation := 0; ; minLocation++ {
		possibleSeedNum := minLocation
		for _, reverter := range *reverters {
			possibleSeedNum = reverter.Revert(possibleSeedNum)
		}
		if slices.Contains(*seeds, possibleSeedNum) {
			return minLocation
		}
	}
}

func solvePartTwo(seeds *[]int, reverters *[]Reverter) int {
	for minLocation := 0; ; minLocation++ {
		possibleSeedNum := minLocation
		for _, reverter := range *reverters {
			possibleSeedNum = reverter.Revert(possibleSeedNum)
		}
		for i := 0; i < len(*seeds); i += 2 {
			seedRangeStart := (*seeds)[i]
			rangeLength := (*seeds)[i+1]
			relativePossibleSeed := possibleSeedNum - seedRangeStart
			if relativePossibleSeed >= 0 && relativePossibleSeed < rangeLength {
				return minLocation
			}
		}
	}
}

func parseLineData(lines *[]string) (*[]int, *[]Reverter) {
	seeds := strings.Split(strings.TrimSpace(strings.Split((*lines)[0], "seeds:")[1]), " ")
	var seedsData []int
	for _, seed := range seeds {
		seedNum, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seedsData = append(seedsData, seedNum)
	}

	*lines = (*lines)[3:]
	reverters := []Reverter{}
	reverterLines := []string{}
	re := regexp.MustCompile(`.*-.*-.*\smap:`)
	for _, line := range *lines {
		lineMatches := re.MatchString(line)
		if lineMatches {
			reverters = append(reverters, *Get(&reverterLines))

			reverterLines = []string{}
		} else {
			reverterLines = append(reverterLines, line)
		}
	}
	reverters = append(reverters, *Get(&reverterLines))
	slices.Reverse(reverters)

	return &seedsData, &reverters
}

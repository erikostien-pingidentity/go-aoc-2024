package cubecon

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/erikostien-pingidentity/go-aoc-2024/common"
)

const (
	totalRedCubes   = 12
	totalGreenCubes = 13
	totalBlueCubes  = 14
)

func SolveDayTwo(dataFileName string) (int, int) {
	// Get all lines from data file
	lines := common.GetLinesFromTestdataFile(dataFileName)

	// Use Capturing groups to get the game number (group1) and the game input (group2)
	re := regexp.MustCompile(`^Game\s*(\d*):\s*(.*)$`)

	partOneTotal, partTwoTotal := 0, 0

	// Loop through each dataLine and parse the data
	for _, dataLine := range lines {
		matches := re.FindAllStringSubmatch(dataLine, -1)

		// Extract the game number group value
		gameNumber, err := strconv.Atoi(matches[0][1])
		if err != nil {
			panic(err)
		}

		// Extract the game input group value
		gameInput := matches[0][2]

		// Solve for part one
		if gameIsValid(gameInput) {
			partOneTotal += gameNumber
		}

		// Solve for part two
		partTwoTotal += getPowerOfMinimumSet(gameInput)
	}
	return partOneTotal, partTwoTotal
}

func gameIsValid(gameInput string) bool {
	// Seperate each game into sets
	sets := strings.Split(gameInput, "; ")

	for _, set := range sets {
		// Seperate each set to their color amounts pulled from the bag
		colorsPulled := strings.Split(set, ", ")
		for _, colorPulled := range colorsPulled {
			valueAndColor := strings.Split(colorPulled, " ")
			value, err := strconv.Atoi(valueAndColor[0])
			if err != nil {
				panic(err)
			}
			color := valueAndColor[1]
			switch color {
			case "red":
				if value > totalRedCubes {
					return false
				}
			case "green":
				if value > totalGreenCubes {
					return false
				}
			case "blue":
				if value > totalBlueCubes {
					return false
				}
			default:
				panic("Invalid color parsed")
			}
		}
	}
	return true
}

func getPowerOfMinimumSet(gameInput string) int {
	maxRed, maxBlue, maxGreen := -1, -1, -1

	// Seperate each game into sets
	sets := strings.Split(gameInput, "; ")

	for _, set := range sets {
		// Seperate each set to their color amounts pulled from the bag
		colorsPulled := strings.Split(set, ", ")
		for _, colorPulled := range colorsPulled {
			valueAndColor := strings.Split(colorPulled, " ")
			value, err := strconv.Atoi(valueAndColor[0])
			if err != nil {
				panic(err)
			}
			color := valueAndColor[1]
			switch color {
			case "red":
				if value > maxRed {
					maxRed = value
				}
			case "green":
				if value > maxGreen {
					maxGreen = value
				}
			case "blue":
				if value > maxBlue {
					maxBlue = value
				}
			default:
				panic("Invalid color parsed")
			}
		}
	}
	return maxRed * maxBlue * maxGreen
}

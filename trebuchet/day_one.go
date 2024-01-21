package trebuchet

import (
	"strconv"
	"strings"

	"github.com/erikostien-pingidentity/go-aoc-2024/common"
)

func SolveDayOne(dataFileName string) (int, int) {
	// Get all lines from data file
	lines := common.GetLinesFromTestdataFile(dataFileName)

	partOneTotal, partTwoTotal := 0, 0

	for _, dataLine := range lines {
		partOneTotal += getCalibrationValue(dataLine, false)

		partTwoTotal += getCalibrationValue(dataLine, true)
	}

	return partOneTotal, partTwoTotal
}

func getCalibrationValue(line string, withReplace bool) int {
	// Create a string replacer to handle word versions of digits
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")

	//find first digit
	firstDigit := ""
	substring := ""
	for _, char := range line {
		if strings.Contains("123456789", string(char)) {
			firstDigit = string(char)
			break
		}

		if !withReplace {
			continue
		}

		substring += string(char)
		replaced := replacer.Replace(substring)
		if len(replaced) != len(substring) {
			firstDigit = string(replaced[len(replaced)-1])
			break
		}
	}

	//find last digit
	lastDigit := ""
	substring = ""
	for i := range line {
		if strings.Contains("123456789", string(line[len(line)-i-1])) {
			lastDigit = string(line[len(line)-i-1])
			break
		}

		if !withReplace {
			continue
		}

		substring = string(line[len(line)-i-1]) + substring
		replaced := replacer.Replace(substring)
		if len(replaced) != len(substring) {
			lastDigit = string(replaced[0])
			break
		}
	}

	if firstDigit == "" {
		return 0
	}

	calibrationValue, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		panic(err)
	}

	return calibrationValue
}

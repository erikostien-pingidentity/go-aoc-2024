package trebuchet

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func SolveDayOne(dataFileName string) (int, int) {
	// Get the current path of this file during execution
	_, thisFilePath, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("error: no caller information"))
	}

	thisFileDir := path.Dir(thisFilePath)

	// Get the full pathname for the data
	dataFilePath := fmt.Sprintf("%s/testdata/%s", thisFileDir, dataFileName)

	// Open the puzzle data that implements the io.Reader interface
	dataFile, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}

	// Make sure to close the opened file once this scope is closed
	defer dataFile.Close()

	// Create a new Scanner to enable reading the file line by line
	scanner := bufio.NewScanner(dataFile)

	partOneTotal, partTwoTotal := 0, 0

	for scanner.Scan() {
		dataLine := scanner.Text()

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

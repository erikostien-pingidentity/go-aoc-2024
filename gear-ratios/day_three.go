package gearratios

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/erikostien-pingidentity/go-aoc-2024/common"
)

type Part struct {
	Symbol      rune
	PartNumbers []int
}

func SolveDayThree(dataFilename string) (int, int) {
	lines := common.GetLinesFromTestdataFile(dataFilename)

	padSchematic(&lines)

	parts := parseAllParts(&lines)

	partOneSolution := solvePartOne(&parts)

	partTwoSolution := solvePartTwo(&parts)

	return partOneSolution, partTwoSolution
}

// Sum all part numbers for all parts and return the result
func solvePartOne(parts *[]Part) int {
	total := 0
	for _, part := range *parts {
		for _, partNumber := range part.PartNumbers {
			total += partNumber
		}
	}
	return total
}

// Find the gear ratio of any gear part (symbol: '*') that has exactly two partnumbers attached
// gear ratio = two part numbers multiplied
// Solution is all gear ratios summed together
func solvePartTwo(parts *[]Part) int {
	total := 0
	for _, part := range *parts {
		if part.Symbol == '*' && len(part.PartNumbers) == 2 {
			total += part.PartNumbers[0] * part.PartNumbers[1]
		}
	}
	return total
}

// Loop through the schematic, and parse all part numbers attached to all parts (symbols)
func parseAllParts(lines *[]string) []Part {
	validPartSymbols := "!@#$%^&*-=+/"

	parts := []Part{}
	for i := 1; i < len((*lines))-1; i++ {
		line := (*lines)[i]
		for j, char := range line {
			if strings.ContainsRune(validPartSymbols, char) {
				parts = append(parts, parsePart(lines, char, i, j))
			}
		}
	}

	return parts
}

func parsePart(lines *[]string, char rune, rowIndex int, colIndex int) Part {
	// Set up regex statements
	topBottomRe := regexp.MustCompile(`^(\d{3})|(\d{3})$|^.+(\d{3}).+$|^\.(\d{2})|(\d{2})\.$|^.{2,}(\d{2}).{2,}$|^\.{2}(\d)|(\d)\.{2}$|^.{3,}(\d).{3,}$`)
	leftMiddleRe := regexp.MustCompile(`(\d{1,3})$`)
	rightMiddleRe := regexp.MustCompile(`^(\d{1,3})`)

	attachedPartNumbers := []string{}

	//Parse Top
	matches := topBottomRe.FindAllStringSubmatch((*lines)[rowIndex-1][colIndex-3:colIndex+4], -1)
	for _, match := range matches {
		for i := 1; i < len(match); i++ {
			if match[i] != "" {
				partNumber := strings.ReplaceAll(match[i], ".", "")
				attachedPartNumbers = append(attachedPartNumbers, partNumber)
				break
			}
		}
	}

	//Parse Left
	leftPartNumbers := leftMiddleRe.FindAllString((*lines)[rowIndex][colIndex-3:colIndex], -1)
	attachedPartNumbers = append(attachedPartNumbers, leftPartNumbers...)

	//Parse Right
	rightPartNumbers := rightMiddleRe.FindAllString((*lines)[rowIndex][colIndex+1:colIndex+4], -1)
	attachedPartNumbers = append(attachedPartNumbers, rightPartNumbers...)

	//Parse Bottom
	matches = topBottomRe.FindAllStringSubmatch((*lines)[rowIndex+1][colIndex-3:colIndex+4], -1)
	for _, match := range matches {
		for i := 1; i < len(match); i++ {
			if match[i] != "" {
				partNumber := strings.ReplaceAll(match[i], ".", "")
				attachedPartNumbers = append(attachedPartNumbers, partNumber)
				break
			}
		}
	}

	//String to int convert all part numbers
	intPartNumbers := []int{}
	for _, stringPartNumber := range attachedPartNumbers {
		intPartNumber, err := strconv.Atoi(stringPartNumber)
		if err != nil {
			panic(err)
		}
		intPartNumbers = append(intPartNumbers, intPartNumber)
	}
	resultPart := Part{
		Symbol:      char,
		PartNumbers: intPartNumbers,
	}

	return resultPart
}

// Pad the schematic to avoid out-of-bounds logic nightmare
func padSchematic(lines *[]string) {
	dataLineLength := len((*lines)[0])

	// Pad the schematic sides with empty '.' data
	for index, line := range *lines {
		(*lines)[index] = fmt.Sprintf("...%s...", line)
	}

	// Pad the schematic top and bottom with empty data lines
	emptyDataLine := strings.Repeat(".", dataLineLength+6)
	*lines = common.Prepend(*lines, emptyDataLine)
	*lines = append(*lines, emptyDataLine)
}

package main

import (
	"fmt"

	"github.com/erikostien-pingidentity/go-aoc-2024/trebuchet"
)

func main() {
	// Solve Day One
	answerPartOne, answerPartTwo := trebuchet.SolveDayOne("puzzle_data")
	fmt.Printf("Answer to Day One Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day One Part Two: %d\n", answerPartTwo)
}

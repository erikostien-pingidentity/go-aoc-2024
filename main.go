package main

import (
	"fmt"

	cubecon "github.com/erikostien-pingidentity/go-aoc-2024/cube-conundrum"
	"github.com/erikostien-pingidentity/go-aoc-2024/trebuchet"
)

func main() {
	// Solve Day One
	answerPartOne, answerPartTwo := trebuchet.SolveDayOne("puzzle_data")
	fmt.Printf("Answer to Day One Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day One Part Two: %d\n", answerPartTwo)

	// Solve Day Two
	answerPartOne, answerPartTwo = cubecon.SolveDayTwo("puzzle_data")
	fmt.Printf("Answer to Day One Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day One Part Two: %d\n", answerPartTwo)
}

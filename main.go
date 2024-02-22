package main

import (
	"fmt"

	cubecon "github.com/erikostien-pingidentity/go-aoc-2024/cube-conundrum"
	"github.com/erikostien-pingidentity/go-aoc-2024/fertilizer"
	gearratios "github.com/erikostien-pingidentity/go-aoc-2024/gear-ratios"
	"github.com/erikostien-pingidentity/go-aoc-2024/scratchcards"
	"github.com/erikostien-pingidentity/go-aoc-2024/trebuchet"
)

func main() {
	// Solve Day One
	answerPartOne, answerPartTwo := trebuchet.SolveDayOne("puzzle_data")
	fmt.Printf("Answer to Day One Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day One Part Two: %d\n", answerPartTwo)

	// Solve Day Two
	answerPartOne, answerPartTwo = cubecon.SolveDayTwo("puzzle_data")
	fmt.Printf("Answer to Day Two Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day Two Part Two: %d\n", answerPartTwo)

	// Solve Day Three
	answerPartOne, answerPartTwo = gearratios.SolveDayThree("puzzle_data")
	fmt.Printf("Answer to Day Three Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day Three Part Two: %d\n", answerPartTwo)

	// Solve Day Four
	answerPartOne, answerPartTwo = scratchcards.SolveDayFour("puzzle_data")
	fmt.Printf("Answer to Day Four Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day Four Part Two: %d\n", answerPartTwo)

	// Solve Day Five
	answerPartOne, answerPartTwo = fertilizer.SolveDayFive("puzzle_data")
	fmt.Printf("Answer to Day Five Part One: %d\n", answerPartOne)
	fmt.Printf("Answer to Day Five Part Two: %d\n", answerPartTwo)
}

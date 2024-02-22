package scratchcards

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/erikostien-pingidentity/go-aoc-2024/common"
)

type ScratchCard struct {
	WinningNumbers []string
	GivenNumbers   []string
	CardCopies     int
	Matches        int
}

func SolveDayFour(dataFilename string) (int, int) {
	lines := common.GetLinesFromTestdataFile(dataFilename)

	scratchCards := ParseScratchCardData(&lines)

	CalculateScratchCardMatches(scratchCards)

	partOneAnswer := solvePartOne(scratchCards)

	partTwoAnswer := SolvePartTwo(scratchCards)

	return partOneAnswer, partTwoAnswer
}

func ParseScratchCardData(lines *[]string) *map[int]*ScratchCard {
	scratchCards := map[int]*ScratchCard{}

	for _, line := range *lines {
		re := regexp.MustCompile(`\s+`)

		lineData := strings.Split(line, ":")
		cardNumber, err := strconv.Atoi(re.Split(lineData[0], -1)[1])
		if err != nil {
			panic(err)
		}

		lineData = strings.Split(strings.TrimSpace(lineData[1]), "|")
		winningNumbers := re.Split(strings.TrimSpace(lineData[0]), -1)
		givenNumbers := re.Split(strings.TrimSpace(lineData[1]), -1)

		parsedScratchCard := ScratchCard{
			WinningNumbers: winningNumbers,
			GivenNumbers:   givenNumbers,
			CardCopies:     1,
			Matches:        0,
		}

		scratchCards[cardNumber] = &parsedScratchCard
	}

	return &scratchCards
}

func CalculateScratchCardMatches(scratchCards *map[int]*ScratchCard) {
	for _, card := range *scratchCards {
		for _, givenNumber := range card.GivenNumbers {
			if slices.Contains(card.WinningNumbers, givenNumber) {
				card.Matches += 1
			}
		}
	}
}

func solvePartOne(scratchCards *map[int]*ScratchCard) int {
	totalPoints := 0

	for _, card := range *scratchCards {
		cardPoints := 0
		for i := 0; i < card.Matches; i++ {
			if cardPoints == 0 {
				cardPoints = 1
			} else {
				cardPoints *= 2
			}
		}
		totalPoints += cardPoints
	}

	return totalPoints
}

func SolvePartTwo(scratchCards *map[int]*ScratchCard) int {
	for cardNumber := 1; cardNumber <= len(*scratchCards); cardNumber++ {
		card := (*scratchCards)[cardNumber]
		for i := 1; i <= card.Matches; i++ {
			relativeCard := (*scratchCards)[cardNumber+i]
			relativeCard.CardCopies += card.CardCopies
		}
	}

	totalCards := 0

	for _, card := range *scratchCards {
		totalCards += card.CardCopies
	}

	return totalCards
}

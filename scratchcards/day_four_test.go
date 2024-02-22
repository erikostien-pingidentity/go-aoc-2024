package scratchcards_test

import (
	"testing"

	"github.com/erikostien-pingidentity/go-aoc-2024/scratchcards"
)

func TestParseScratchCardData(t *testing.T) {
	testLines := []string{
		"Card 1:  1  2  3 |  4  5  6",
		"Card 2: 10 | 11",
	}

	scratchCards := scratchcards.ParseScratchCardData(&testLines)

	if len(*scratchCards) != 2 {
		t.Errorf("Expected 2 parsed cards, got %d", len(*scratchCards))
	}

	cardOne := (*scratchCards)[1]
	cardTwo := (*scratchCards)[2]

	if len(cardOne.WinningNumbers) != 3 {
		t.Errorf("Expected 3 winning numbers, got %d", len(cardOne.WinningNumbers))
	}

	if len(cardOne.GivenNumbers) != 3 {
		t.Errorf("Expected 3 given numbers, got %d", len(cardOne.GivenNumbers))
	}

	if len(cardTwo.WinningNumbers) != 1 {
		t.Errorf("Expected 1 winning number, got %d", len(cardTwo.WinningNumbers))
	}

	if len(cardTwo.GivenNumbers) != 1 {
		t.Errorf("Expected 1 given number, got %d", len(cardTwo.GivenNumbers))
	}

	if cardOne.CardCopies != 1 {
		t.Errorf("Expected 1 card copy be default, got %d", cardOne.CardCopies)
	}

	if cardOne.Matches != 0 {
		t.Errorf("Expected 0 matches be default, got %d", cardOne.Matches)
	}
}

func TestCalculateScratchCardMatches(t *testing.T) {
	testLines := []string{
		"Card 1:  1  2  3 |  4  5  6",
		"Card 2: 10 | 11 10 10 10  1",
	}

	scratchCards := scratchcards.ParseScratchCardData(&testLines)

	scratchcards.CalculateScratchCardMatches(scratchCards)

	cardOne := (*scratchCards)[1]
	cardTwo := (*scratchCards)[2]

	if cardOne.Matches != 0 {
		t.Errorf("Expected 0 matches, got %d", cardOne.Matches)
	}

	if cardTwo.Matches != 3 {
		t.Errorf("Expected 3 matches, got %d", cardTwo.Matches)
	}
}

func TestSolveDayFour(t *testing.T) {
	expectedDayOne, expectedDayTwo := 13, 30

	dayOneAnswer, dayTwoAnswer := scratchcards.SolveDayFour("example_data")

	if dayOneAnswer != expectedDayOne {
		t.Errorf("Expected %d points, got %d", expectedDayOne, dayOneAnswer)
	}

	if dayTwoAnswer != expectedDayTwo {
		t.Errorf("Expected %d points, got %d", expectedDayTwo, dayTwoAnswer)
	}
}

func TestCustom(t *testing.T) {
	testLines := []string{
		"Card   1:  5 27 94 20 50  7 98 41 67 34 | 34  9 20 90  7 77 44 71 27 12 98  1 79 96 24 51 25 84 67 41  5 13 78 31 26",
		"Card   2: 52 14 37 45 82 39 73 67 72 90 | 72 78 37 25 39 68 23 45 73 90 86  2 85 57 80 62 22 26 92 67 82 95 66 14 52",
		"Card   3: 33 86 80 53 45 32 25 29 84 89 | 89 53 74 49 73 30 25 82 84 44 59 86 32 54 96 41 26 80 68 29 45 33 23 99 17",
		"Card   4: 64 25  5  1 46 75 45 55 21  7 | 93 62 21 60 46 44 96 88 12 63 85 91 14 55 68 67 16 74 45 41 75 70 25 36 78",
		"Card   5:  1 |  2",
		"Card   6:  1 |  2",
		"Card   7:  1 |  2",
		"Card   8:  1 |  2",
		"Card   9:  1 |  2",
		"Card  10:  1 |  2",
		"Card  11:  1 |  2",
		"Card  12:  1 |  2",
		"Card  13:  1 |  2",
		"Card  14:  1 |  2",
	}
	expectedCardMatches := []int{8, 10, 10, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedCardCopies := []int{1, 2, 4, 8, 16, 16, 16, 16, 16, 15, 7, 7, 5, 1}

	scratchCards := scratchcards.ParseScratchCardData(&testLines)

	scratchcards.CalculateScratchCardMatches(scratchCards)

	answerPartTwo := scratchcards.SolvePartTwo(scratchCards)

	if answerPartTwo != 130 {
		t.Errorf("Expected 115 total card copies, got %d", answerPartTwo)
	}

	for i := 1; i < len(*scratchCards); i++ {
		card := (*scratchCards)[i]

		if card.Matches != expectedCardMatches[i-1] {
			t.Errorf("Expected %d matches, got %d", expectedCardMatches[i-1], card.Matches)
		}

		if card.CardCopies != expectedCardCopies[i-1] {
			t.Errorf("Expected %d matches, got %d", expectedCardCopies[i-1], card.CardCopies)
		}
	}
}

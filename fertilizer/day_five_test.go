package fertilizer

import "testing"

func TestSolveDayFive(t *testing.T) {
	answerPartOne, answerPartTwo := SolveDayFive("example_data")

	expectedPartOne, expectedPartTwo := 35, 46

	if answerPartOne != expectedPartOne {
		t.Errorf("Expected lowest location %d, got %d", expectedPartOne, answerPartOne)
	}

	if answerPartTwo != expectedPartTwo {
		t.Errorf("Expected lowest location %d, got %d", expectedPartTwo, answerPartTwo)
	}
}

package cubecon

import "testing"

func TestSolve(t *testing.T) {
	partOneResult, partTwoResult := SolveDayTwo("example_data")
	partOneExpected, partTwoExpected := 8, 2286

	if partOneResult != partOneExpected {
		t.Errorf("Expected %d, got %d", partOneExpected, partOneResult)
	}

	if partTwoResult != partTwoExpected {
		t.Errorf("Expected %d, got %d", partTwoExpected, partTwoResult)
	}
}

func TestGameIsValid(t *testing.T) {
	games := []string{
		"4 green; 1 red, 3 blue",
		"50 red",
		"5 green, 1 blue, 2 red; 40 green",
		"5 green, 1 blue, 2 red; 5 green, 1 blue, 2 red; 5 green, 1 blue, 2 red; 5 green, 1 blue, 2 red",
	}

	expecteds := []bool{
		true,
		false,
		false,
		true,
	}

	for i, game := range games {
		result := gameIsValid(game)
		expected := expecteds[i]

		if expected != result {
			t.Errorf("Expected %t, got %t", expected, result)
		}
	}
}

func TestGetPowerOfMinimumSet(t *testing.T) {
	games := []string{
		"4 green; 1 red, 3 blue",
		"42 green, 1 blue, 2 red; 40 green",
		"5 green, 1 blue, 2 red; 5 green, 1 blue, 2 red; 5 green, 1 blue, 2 red; 5 green, 1 blue, 2 red",
		"1 green; 1 blue; 1 red",
		"5 green, 5 blue, 5 red; 1 green; 1 blue; 1 red",
	}

	expecteds := []int{
		12,
		84,
		10,
		1,
		125,
	}

	for i, game := range games {
		result := getPowerOfMinimumSet(game)
		expected := expecteds[i]

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	}
}

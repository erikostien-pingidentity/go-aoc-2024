package trebuchet

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	partOneResult, partTwoResult := SolveDayOne("ex_data_pt_1")

	partOneExpected, partTwoExpected := 142, 142

	if partOneResult != partOneExpected {
		t.Errorf("Expected %d, got %d", partOneExpected, partOneResult)
	}

	if partTwoResult != partTwoExpected {
		t.Errorf("Expected %d, got %d", partTwoExpected, partTwoResult)
	}
}

func TestSolvePartTwo(t *testing.T) {
	partOneResult, partTwoResult := SolveDayOne("ex_data_pt_2")

	partOneExpected, partTwoExpected := 209, 281

	if partOneResult != partOneExpected {
		t.Errorf("Expected %d, got %d", partOneExpected, partOneResult)
	}

	if partTwoResult != partTwoExpected {
		t.Errorf("Expected %d, got %d", partTwoExpected, partTwoResult)
	}
}

func TestGetCalibrationValue(t *testing.T) {
	testStrings := []string{
		"kljsdf56kljdf",
		"1lksjdg5",
		"kjddg7sdklghskld",
		"one5nine",
		"eight7",
		"twone",
	}

	expectedInts := []int{
		56,
		15,
		77,
		19,
		87,
		21,
	}

	for index, value := range testStrings {
		result := getCalibrationValue(value, true)
		if result != expectedInts[index] {
			t.Errorf("Expected %d, got %d", expectedInts[index], result)
		}
	}
}

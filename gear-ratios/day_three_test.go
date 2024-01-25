package gearratios

import (
	"reflect"
	"testing"
)

func TestPadSchematic(t *testing.T) {
	testSchem := []string{"."}

	expectedSchemLine := "......."

	padSchematic(&testSchem)

	if len(testSchem) != 3 {
		t.Errorf("Expected Schematic Size %d, got %d", 3, len(testSchem))
	}

	for _, line := range testSchem {
		if line != expectedSchemLine {
			t.Errorf("Expected Schematic Line %q, got %q", expectedSchemLine, line)
		}
	}
}

func TestParseParEmpty(t *testing.T) {
	testSchem := []string{
		".......",
		"...=...",
		".......",
	}

	expectedPart := Part{
		Symbol:      '=',
		PartNumbers: []int{},
	}

	resultPart := parsePart(&testSchem, '=', 1, 3)

	if !reflect.DeepEqual(expectedPart, resultPart) {
		t.Errorf("Expected part %q, got %q", expectedPart, resultPart)
	}
}

func TestParsePartPopulated(t *testing.T) {
	testSchem := []string{
		"..5.123",
		".98=97.",
		"23.124.",
	}

	expectedPart := Part{
		Symbol: '=',
		PartNumbers: []int{
			5,
			123,
			98,
			97,
			124,
		},
	}

	resultPart := parsePart(&testSchem, '=', 1, 3)

	if !reflect.DeepEqual(expectedPart, resultPart) {
		t.Errorf("Expected part %q, got %q", expectedPart, resultPart)
	}
}

func TestParseAllParts(t *testing.T) {
	testSchem := []string{
		"..5.123",
		"...=...",
		"23..24.",
		".......",
		"4..*...",
		".123.45",
	}

	expectedParts := []Part{
		{
			Symbol: '=',
			PartNumbers: []int{
				5,
				123,
				24,
			},
		},
		{
			Symbol: '*',
			PartNumbers: []int{
				123,
			},
		},
	}

	resultParts := parseAllParts(&testSchem)

	if !reflect.DeepEqual(expectedParts, resultParts) {
		t.Errorf("Expected parts %q, got %q", expectedParts, resultParts)
	}

}

func TestSolveDayThree(t *testing.T) {
	expectedPartOne, expectedPartTwo := 4361, 467835

	resultPartOne, resultPartTwo := SolveDayThree("example_data")

	if resultPartOne != expectedPartOne {
		t.Errorf("Part One: Expected %d, got %d", expectedPartOne, resultPartOne)
	}

	if resultPartTwo != expectedPartTwo {
		t.Errorf("Part Two: Expected %d, got %d", expectedPartTwo, resultPartTwo)
	}
}

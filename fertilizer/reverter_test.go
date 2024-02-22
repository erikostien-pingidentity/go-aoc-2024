package fertilizer

import "testing"

func TestConverterGet(t *testing.T) {
	testLines := []string{
		"50 98 2",
		"52 50 48",
		"",
	}

	reverter := Get(&testLines)

	expectedReversion := 98
	actualReversion := reverter.Revert(50)

	if actualReversion != expectedReversion {
		t.Errorf("Expected conversion to be %d, got %d", expectedReversion, actualReversion)
	}

	expectedReversion = 53
	actualReversion = reverter.Revert(55)

	if actualReversion != expectedReversion {
		t.Errorf("Expected conversion to be %d, got %d", expectedReversion, actualReversion)
	}

	expectedReversion = 10
	actualReversion = reverter.Revert(10)

	if actualReversion != expectedReversion {
		t.Errorf("Expected conversion to be %d, got %d", expectedReversion, actualReversion)
	}
}

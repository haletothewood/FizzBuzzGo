package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	resultsString := FizzBuzz(100)
	expectedString := ""
	if resultsString != expectedString {
		t.Error("Expected: ", resultsString, ", got: ", expectedString)
	}
}

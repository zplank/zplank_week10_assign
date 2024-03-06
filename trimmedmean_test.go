// trimmedmean_test.go

package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	// create a slice of 100 integers
	intTestData := make([]float64, 100)
	for i := range intTestData {
		intTestData[i] = float64(i + 1)
	}

	// create a slice of 100 floats
	floatTestData := make([]float64, 100)
	for i := range floatTestData {
		floatTestData[i] = float64(i)*0.5 + 0.5
	}

	// mean(x.trim = 0.05)
	lowerTrim, upperTrim := 0.05, 0.05

	// expected means
	expectedMeanInts := 50.5
	expectedMeanFloats := 25.25

	resultInts, err := TrimmedMean(intTestData, lowerTrim, upperTrim)
	if err != nil {
		t.Fatalf("TrimmedMean returned an error for integers: %v", err)
	}

	resultFloats, err := TrimmedMean(floatTestData, lowerTrim, upperTrim)
	if err != nil {
		t.Fatalf("TrimmedMean returned an error for floats: %v", err)
	}

	if resultInts != expectedMeanInts {
		t.Errorf("Expected mean for integers: %v, got: %v", expectedMeanInts, resultInts)
	}

	if resultFloats != expectedMeanFloats {
		t.Errorf("Expected mean for floats: %v, got: %v", expectedMeanFloats, resultFloats)
	}
}

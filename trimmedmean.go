package trimmedmean

import (
	"errors"
	"sort"
)

// calculate the trimmed mean of a slice of float64 by timming a proportion (0.05) of high and low before calculating mean
func TrimmedMean(data []float64, lowerTrim, upperTrim float64) (float64, error) {
	if lowerTrim < 0 || lowerTrim > 0.05 || upperTrim < 0 || upperTrim > 0.05 {
		return 0, errors.New("trim proportions must be between 0 and 0.05")
	}
	if upperTrim == 0 {
		upperTrim = lowerTrim
	}

	n := len(data)
	if n == 0 {
		return 0, errors.New("slice is empty")
	}

	sort.Float64s(data)

	lowerCount := int(float64(n) * lowerTrim)
	upperCount := int(float64(n) * upperTrim)
	if n-lowerCount-upperCount <= 0 {
		return 0, errors.New("trimming proportions too large")
	}

	trimmedData := data[lowerCount : n-upperCount]
	sum := 0.0
	for _, v := range trimmedData {
		sum += v
	}
	trimmedMean := sum / float64(len(trimmedData))
	return trimmedMean, nil
}

package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(stringValues []string) ([]float64, error) {
	var floatValues []float64

	for _, str := range stringValues {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float: " + str)
		}
		floatValues = append(floatValues, val)
	}

	return floatValues, nil
}

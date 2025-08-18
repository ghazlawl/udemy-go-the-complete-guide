package io

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	var valueString = fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueString), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0.0, errors.New("failed to read file")
	}

	valueString := string(data)

	value, err := strconv.ParseFloat(valueString, 64)

	if err != nil {
		return 0.0, errors.New("failed to parse value from file")
	}

	return value, nil
}

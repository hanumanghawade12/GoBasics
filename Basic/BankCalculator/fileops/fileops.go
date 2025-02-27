package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	// write to file
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	// read from file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return 1000, errors.New("error reading file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println("Error converting balance to float")
		return 1000, errors.New("error converting balance to float")
	}
	return value, nil
}

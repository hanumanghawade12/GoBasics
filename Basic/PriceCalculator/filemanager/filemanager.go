package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	inputPath  string
	outputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.inputPath)

	if err != nil {
		return nil, errors.New("error opening file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("error reading file")
		file.Close()
		return nil, errors.New("error reading file")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.outputPath)
	if err != nil {
		fmt.Println("error creating file")
		return errors.New("error creating file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("error encoding json")
		file.Close()
		return errors.New("error encoding json")
	}
	file.Close()
	return nil
}

func NewFileManager(inputPath, outputPath string) FileManager {
	return FileManager{
		inputPath:  inputPath,
		outputPath: outputPath,
	}
}

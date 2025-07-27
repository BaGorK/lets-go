package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("error opening file: " + fm.InputFilePath)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		err := file.Close()
		if err != nil {
			return nil, errors.New("error closing file: " + fm.InputFilePath)
		}
		return nil, errors.New("error reading file: " + fm.InputFilePath)
	}

	err = file.Close()
	if err != nil {
		return nil, errors.New("error closing file: " + fm.InputFilePath)
	}

	return lines, nil
}

func (fm Filemanager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("error creating file: " + fm.OutputFilePath)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		errClose := file.Close()
		if errClose != nil {
			return errors.New("error closing file: " + fm.OutputFilePath)
		}
		return errors.New("error writing to file: " + fm.OutputFilePath)
	}

	err = file.Close()
	if err != nil {
		return errors.New("error closing file: " + fm.OutputFilePath)
	}

	return nil
}

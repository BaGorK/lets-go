package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error opening file: " + path)
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
			return nil, errors.New("error closing file: " + path)
		}
		return nil, errors.New("error reading file: " + path)
	}

	err = file.Close()
	if err != nil {
		return nil, errors.New("error closing file: " + path)
	}

	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("error creating file: " + path)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		errClose := file.Close()
		if errClose != nil {
			return errors.New("error closing file: " + path)
		}
		return errors.New("error writing to file: " + path)
	}

	err = file.Close()
	if err != nil {
		return errors.New("error closing file: " + path)
	}

	return nil
}

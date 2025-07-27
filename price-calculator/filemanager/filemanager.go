package filemanager

import (
	"bufio"
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

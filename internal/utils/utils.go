package utils

import (
	"bufio"
	"log"
	"os"
)

func ProcessInput(filepath string, onLine func(line string) error) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := onLine(scanner.Text()); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func GetAllInputLines(filepath string) ([]string, error) {
	var inputLines []string

	err := ProcessInput(filepath, func(line string) error {
		inputLines = append(inputLines, line)
		return nil
	})

	return inputLines, err
}

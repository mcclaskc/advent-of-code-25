package utils

import (
	"bufio"
	"fmt"
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
		fmt.Println(line)
		inputLines = append(inputLines, line)
		return nil
	})

	return inputLines, err
}

func Debug(prefix string, printable any) {
	//if os.Getenv("DEBUG") == "TRUE" {
	println("**************************")
	println(prefix)
	println("--------------------------")
	fmt.Println(printable)
	println("--------------------------")
	//}
}

const (
	AnsiColorReset   = "\033[0m"
	AnsiColorRed     = "\033[31m"
	AnsiColorGreen   = "\033[32m"
	AnsiColorYellow  = "\033[33m"
	AnsiColorBlue    = "\033[34m"
	AnsiColorMagenta = "\033[35m"
	AnsiColorCyan    = "\033[36m"
	AnsiColorWhite   = "\033[37m"
)

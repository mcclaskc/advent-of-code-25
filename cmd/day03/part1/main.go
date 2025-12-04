package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

// 987654321111111 -> 98
func maximizeJolts(batteryBank string) int {
	chars := strings.Split(batteryBank, "")
	numChars := len(chars)
	firstBiggest, _ := strconv.Atoi(chars[0])
	lastDigit, _ := strconv.Atoi(chars[numChars-1])
	secondBiggest := lastDigit
	for i := 1; i < numChars-1; i++ {
		digit, _ := strconv.Atoi(chars[i])
		if digit > firstBiggest {
			firstBiggest = digit
			secondBiggest = lastDigit
		} else if digit > secondBiggest {
			secondBiggest = digit
		}
	}
	return firstBiggest*10 + secondBiggest
}

func main() {
	sum := 0
	utils.ProcessInput("../input.txt", func(line string) error {
		sum += maximizeJolts(line)
		return nil
	})
	fmt.Printf("Final Sum: %d", sum)
}

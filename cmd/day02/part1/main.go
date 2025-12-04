package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func isInvalid(id string) bool {
	// if it's odd, we know it can't be invalid, since it won't be 'two repeated numbers'
	numDigits := len(id)
	if id == "" || numDigits%2 != 0 {
		return false
	}
	midIndex := numDigits / 2
	firstHalf := id[:midIndex]
	secondHalf := id[midIndex:]
	return firstHalf == secondHalf
}

// 111-222
func findInvalidIdsInRange(rangeStr string) []int {
	firstStr, lastStr, _ := strings.Cut(rangeStr, "-")
	firstInt, _ := strconv.Atoi(firstStr)
	lastInt, _ := strconv.Atoi(lastStr)

	invalidIds := []int{}
	// check first manually to avoid double stringifying
	if isInvalid(firstStr) {
		invalidIds = append(invalidIds, firstInt)
	}
	for i := firstInt + 1; i < lastInt; i++ {
		iStr := strconv.Itoa(i)
		if isInvalid(iStr) {
			invalidIds = append(invalidIds, i)
		}
	}
	// check last manually to avoid double stringifying
	if isInvalid(lastStr) {
		invalidIds = append(invalidIds, lastInt)
	}
	return invalidIds
}

func sumAllIds(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}

func main() {
	sum := 0
	utils.ProcessInput("../input.txt", func(line string) error {
		invalidIds := findInvalidIdsInRange(line)
		sum += sumAllIds(invalidIds)
		return nil
	})
	fmt.Printf("Final Sum: %d", sum)
}

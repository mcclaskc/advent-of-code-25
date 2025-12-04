package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func isInvalid(id string) bool {
	numDigits := len(id)
	for patternLength := 1; patternLength <= numDigits/2; patternLength++ {
		if numDigits%patternLength == 0 {
			// split the id string into equal parts of patternLength
			numParts := numDigits / patternLength
			if numParts < 2 {
				continue
			}
			firstPart := id[:patternLength]
			allSame := true
			for i := 1; i < numParts; i++ {
				part := id[i*patternLength : (i+1)*patternLength]
				if part != firstPart {
					allSame = false
					break
				}
			}
			if allSame {
				return true
			}
		}
	}
	return false
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

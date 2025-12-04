package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

const numDigitsNeeded = 12

func maximizeJolts(batteryBank string) int {
	chars := strings.Split(batteryBank, "")
	numChars := len(chars)
	buffer := numChars - numDigitsNeeded
	maxJoltList := make([]string, numDigitsNeeded)
	ci := 0
	for ji := range maxJoltList {
		bi := 0
		fmt.Printf("ji: %d, ci: %d, bi: %d, buffer: %d\n", ji, ci, bi, buffer)
		largestCandidateIndex := ci
		largestCandidate := chars[largestCandidateIndex]
		for bi <= buffer {
			candidateIndex := ci + bi
			fmt.Printf("ji: %d, ci: %d, bi: %d, candidateIndex: %d\n", ji, ci, bi, candidateIndex)
			candidate := chars[candidateIndex]
			if candidate > largestCandidate {
				largestCandidate = candidate
				largestCandidateIndex = candidateIndex
			}
			bi++
		}
		delta := largestCandidateIndex - ci
		buffer -= delta
		fmt.Printf("delta: %d, largestCandidateIndex: %d, ci: %d\n", delta, largestCandidateIndex, ci)
		ci += delta + 1
		maxJoltList[ji] = largestCandidate
	}

	// }
	final, _ := strconv.Atoi(strings.Join(maxJoltList, ""))
	return final
}

func main() {
	sum := 0
	utils.ProcessInput("../input.txt", func(line string) error {
		sum += maximizeJolts(line)
		return nil
	})
	fmt.Printf("Final Sum: %d", sum)
}

package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strings"
)

func splitCount(filePath string) int {
	r := 0
	prevTimelines := []int{}
	utils.ProcessInput(filePath, func(line string) error {
		nextTimelines := []int{}
		currentRow := strings.Split(line, "")
		if r%2 == 0 {
			fmt.Println(currentRow)
			for c, char := range line {
				if r == 0 {
					if char == 'S' {
						nextTimelines = append(nextTimelines, 1)
					} else {
						nextTimelines = append(nextTimelines, 0)
					}
				} else if char == '.' {
					nextTimelines = append(nextTimelines, prevTimelines[c])
					if c > 0 && currentRow[c-1] == "^" {
						// println("above left is splitter")
						nextTimelines[c] += prevTimelines[c-1]
					}
					if c < len(currentRow)-1 && currentRow[c+1] == "^" {
						// println("above left is splitter")
						nextTimelines[c] += prevTimelines[c+1]
					}
				} else {
					nextTimelines = append(nextTimelines, 0)
				}
			}
			// fmt.Printf("nextTimelines: %d\n", nextTimelines)
			prevTimelines = nextTimelines
		} else {
			fmt.Println(prevTimelines)
		}
		r++
		return nil
	})
	// fmt.Printf("FINAL SUMMING: %d\n", prevTimelines)
	sum := 0
	for _, tl := range prevTimelines {
		sum += tl
	}
	return sum
}

func main() {
	result := splitCount("../input.txt")
	fmt.Print("RESULT ---\n")
	fmt.Printf("RESULT: %d\n", result)
	fmt.Print("RESULT ---\n")
}

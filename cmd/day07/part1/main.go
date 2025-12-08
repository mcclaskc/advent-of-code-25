package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
)

func splitCount(filePath string) int {
	sum := 0
	r := 0
	prevBeamLocations := []rune{}
	nextBeamLocations := []rune{}
	utils.ProcessInput(filePath, func(line string) error {
		if r%2 == 0 {
			fmt.Println(line)
			prevCharWasSplitter := false
			for c, char := range line {
				if char == 'S' {
					nextBeamLocations = append(nextBeamLocations, '|')
				} else if char == '^' && prevBeamLocations[c] == '|' {
					// fmt.Printf("%s : b locations\n", string(nextBeamLocations))
					nextBeamLocations[c-1] = '|'
					nextBeamLocations = append(nextBeamLocations, '.')
					nextBeamLocations = append(nextBeamLocations, '|')
					prevCharWasSplitter = true
					// fmt.Printf("%s : a locations\n", string(nextBeamLocations))
					sum++
				} else {
					if !prevCharWasSplitter {
						var nextChar rune
						if c < len(prevBeamLocations) {
							nextChar = prevBeamLocations[c]
						} else {
							nextChar = '.'
						}
						nextBeamLocations = append(nextBeamLocations, nextChar)
					}
					prevCharWasSplitter = false
				}
			}
			prevBeamLocations = nextBeamLocations
			nextBeamLocations = []rune{}
		} else {
			fmt.Println(string(prevBeamLocations))
		}
		r++
		return nil
	})
	return sum
}

func main() {
	result := splitCount("../input.txt")
	fmt.Print("RESULT ---\n")
	fmt.Printf("RESULT: %d\n", result)
	fmt.Print("RESULT ---\n")
}

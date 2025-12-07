package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

func countFresh(filePath string) int {
	sum := 0
	freshRanges := []Range{}
	appendingRanges := true
	utils.ProcessInput(filePath, func(line string) error {
		fmt.Println(appendingRanges)
		fmt.Println(line)
		if line == "" {
			appendingRanges = false
		} else if appendingRanges {
			strs := strings.Split(line, "-")
			min, _ := strconv.Atoi(strs[0])
			max, _ := strconv.Atoi(strs[1])
			freshRanges = append(freshRanges, Range{min, max})
		} else {
			id, _ := strconv.Atoi(line)
			for i := range freshRanges {
				if id >= freshRanges[i].min && id <= freshRanges[i].max {
					sum++
					break
				}
			}
		}
		return nil
	})
	return sum
}

func main() {
	sum := countFresh("../input.txt")
	fmt.Printf("Final Sum: %d", sum)
}

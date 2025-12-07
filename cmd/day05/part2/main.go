package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"sort"
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
		fmt.Println(line)
		if line == "" {
			appendingRanges = false
		} else if appendingRanges {
			strs := strings.Split(line, "-")
			min, _ := strconv.Atoi(strs[0])
			max, _ := strconv.Atoi(strs[1])
			freshRanges = append(freshRanges, Range{min, max})
		}
		fmt.Println(appendingRanges)
		return nil
	})

	// sort by fresh min
	sort.Slice(freshRanges, func(i int, j int) bool {
		return freshRanges[i].min < freshRanges[j].min
	})

	fmt.Println(freshRanges)
	newRanges := []Range{freshRanges[0]}
	for i := 1; i < len(freshRanges); i++ {
		currentRange := newRanges[len(newRanges)-1]
		nextRange := freshRanges[i]
		if currentRange.max < nextRange.min {
			// no overlap, append
			newRanges = append(newRanges, nextRange)
		} else if currentRange.max >= nextRange.max {
			// current includes next, skip over next
		} else {
			// current and next overlap: set the new current
			newRanges[len(newRanges)-1] = Range{currentRange.min, nextRange.max}
		}
	}

	fmt.Println(newRanges)
	for _, r := range newRanges {
		println(r.min, r.max)
		sum += r.max - r.min + 1
	}
	return sum
}

func main() {
	sum := countFresh("../input.txt")
	fmt.Printf("Final Sum: %d", sum)
}

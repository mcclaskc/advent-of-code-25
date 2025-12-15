package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

// type Rectangle struct {
// 	a    int
// 	b    int
// 	area int
// }

func biggestArea(filePath string) int {
	tiles := [][]int{}
	utils.ProcessInput(filePath, func(line string) error {
		numStrs := strings.Split(line, ",")
		x, _ := strconv.Atoi(numStrs[0])
		y, _ := strconv.Atoi(numStrs[1])
		tiles = append(tiles, []int{x, y})
		return nil
	})
	// rectangles := []Rectangle{}
	largestArea := 0
	for i := 0; i+1 < len(tiles); i++ {
		for j := 1; j < len(tiles); j++ {
			deltaX := tiles[i][0] - tiles[j][0]
			deltaY := tiles[i][1] - tiles[j][1]
			// have to add 1 since the coordinates are inclusive
			length := utils.AbsInt(deltaX) + 1
			height := utils.AbsInt(deltaY) + 1
			area := length * height
			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

func main() {
	result := biggestArea("../input.txt")
	fmt.Print("RESULT ---\n")
	fmt.Printf("RESULT: %d\n", result)
	fmt.Print("RESULT ---\n")
}

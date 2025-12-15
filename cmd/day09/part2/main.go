package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Range struct {
	Start, End int
}

var rowRanges = make(map[int][]Range)

// var runeToColor = map[rune]string{
// 	'R': utils.AnsiColorRed,
// 	'G': utils.AnsiColorWhite,
// 	'.': utils.AnsiColorYellow,
// }

func eligibleRectangle(pointA, pointB Point) bool {
	minY := min(pointA.y, pointB.y)
	maxY := max(pointA.y, pointB.y)
	minX := min(pointA.x, pointB.x)
	maxX := max(pointA.x, pointB.x)
	for row := minY; row <= maxY; row++ {
		// see if x is entirely within any of the ranges, if none, return false
		contained := false
		for _, r := range rowRanges[row] {
			if minX >= r.Start && maxX <= r.End {
				contained = true
				break
			}
		}
		if !contained {
			return false
		}
	}
	return true
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func biggestArea(filePath string) int {
	redPoints := []Point{}
	maxX := 0
	maxY := 0
	minX := math.MaxInt
	minY := math.MaxInt
	// printPoints := func(pointsToHighlight ...Point) {
	// 	if os.Getenv("DEBUG") == "TRUE" {
	// 		println("PONTSMAP --------------------------------------------")
	// 		for y := minY; y < maxY+1; y++ {
	// 			for x := minX; x < maxX+1; x++ {
	// 				candidatePoint := Point{x, y}
	// 				r, found := pointsMap[candidatePoint]
	// 				if !found {
	// 					r = '.'
	// 				}
	// 				color := runeToColor[r]
	// 				if len(pointsToHighlight) > 0 && slices.Contains(pointsToHighlight, candidatePoint) {
	// 					color = utils.AnsiColorCyan
	// 				}
	// 				fmt.Printf("%s%c%s", color, r, utils.AnsiColorReset)
	// 			}
	// 			println()
	// 		}
	// 	}
	// }
	// printPoints()
	utils.ProcessInput(filePath, func(line string) error {
		numStrs := strings.Split(line, ",")
		x, _ := strconv.Atoi(numStrs[0])
		y, _ := strconv.Atoi(numStrs[1])
		maxX = max(x, maxX)
		maxY = max(y, maxY)
		minX = min(x, minX)
		minY = min(y, minY)
		newPoint := Point{x, y}
		redPoints = append(redPoints, newPoint)
		return nil
	})
	numReds := len(redPoints)

	print(2)

	print(3)
	rowWalls := make(map[int][]int)
	for i := range numReds {
		red0 := redPoints[i]
		red1 := Point{}
		if i+1 < len(redPoints) {
			red1 = redPoints[i+1]
		} else {
			red1 = redPoints[0]
		}
		if red0.x == red1.x {
			// vertical line (top is smaller in y choords)
			topY, bottomY := minMax(red0.y, red1.y)
			// add a wall entry for every row in the line
			for y := topY; y < bottomY; y++ {
				rowWalls[y] = append(rowWalls[y], red0.x)
			}
		} else {
			// horizontal line (left is smaller in x choords)
			leftX, rightX := minMax(red0.x, red1.x)
			// add the entire line as a range to be filled for this row (y)
			rowRanges[red0.y] = append(rowRanges[red0.y], Range{leftX, rightX})
		}
	}
	print(4)

	for row, walls := range rowWalls {
		// sort so we know we are going left to right
		slices.Sort(walls)

		// loop over each distinct pair [0,1], [2,3], [3,4], etc
		for i := 0; i+1 < len(walls); i += 2 {
			wall0 := walls[i]
			wall1 := walls[i+1]
			rowRanges[row] = append(rowRanges[row], Range{wall0, wall1})
		}
	}

	print(5)

	print(6)

	print(7)
	largestArea := 0

	for i := 0; i+1 < len(redPoints); i++ {
		for j := i + 1; j < len(redPoints); j++ {
			pointA := redPoints[i]
			pointB := redPoints[j]
			isEligible := eligibleRectangle(pointA, pointB)
			if isEligible {
				deltaX := pointA.x - pointB.x
				deltaY := pointA.y - pointB.y
				// have to add 1 since the coordinates are inclusive
				length := utils.AbsInt(deltaX) + 1
				height := utils.AbsInt(deltaY) + 1
				area := length * height
				if area > largestArea {
					utils.Debug("area", area)
					largestArea = area
				}
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

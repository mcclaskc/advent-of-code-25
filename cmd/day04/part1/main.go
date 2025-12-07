package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"os"
	"strings"
)

func printNeighbor(x int, y int, suffix string) {
	if os.Getenv("PRINTXY") != "" || true {
		fmt.Printf("neighbor: %d, %d | %s\n", x, y, suffix)
	}
}

func countTp(grid [][]string) int {
	sum := 0
	numCols := len(grid[0])
	numRows := len(grid)
	success := [][]int{}
	// for every entry in the grid, count the neighbors
	for y := range numRows {
		for x := range numCols {
			numNeighbors := 0
			fmt.Printf("x: %d; y: %d\n", x, y)
			if grid[y][x] == "@" {
				// loop over all cols 1 behind and 1 ahead
				for xDelta := -1; xDelta <= 1; xDelta++ {
					xNeighbor := x + xDelta
					// only consider in bounds xNeighbors
					if xNeighbor >= 0 && xNeighbor < numCols {
						// loop over all rows 1 above and 1 below
						for yDelta := -1; yDelta <= 1; yDelta++ {
							yNeighbor := y + yDelta
							// only consider in bounds yNeighbors
							if yNeighbor >= 0 && yNeighbor < numRows {
								// skip self
								if !(yNeighbor == y && xNeighbor == x) {
									neighborVal := grid[yNeighbor][xNeighbor]
									printNeighbor(xNeighbor, yNeighbor, neighborVal)
									// check if neighbor is a tp roll
									if neighborVal == "@" {
										numNeighbors++
									}
								} else {
									printNeighbor(x, y, "SELF")
								}
							} else {
								printNeighbor(x, y, "Y OOB")
							}
						}
					} else {
						printNeighbor(xNeighbor, -1, "X OOB")
					}
				}

				if numNeighbors < 4 {
					sum++
					fmt.Printf("SUCCESS: %d, %d\n", x, y)
					success = append(success, []int{x, y})
				}
				fmt.Printf("neighbors: %d, sum: %d\n", numNeighbors, sum)
			} else {
				fmt.Println("skipping .")
			}
		}
	}
	fmt.Print(success)
	return sum
}

func main() {
	grid := [][]string{}
	utils.ProcessInput("../input.txt", func(line string) error {
		grid = append(grid, strings.Split(line, ""))
		return nil
	})
	sum := countTp(grid)
	fmt.Printf("Final Sum: %d", sum)
}

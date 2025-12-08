package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func cephMath(filePath string) int {
	grid := [][]string{}
	utils.ProcessInput(filePath, func(line string) error {
		words := strings.Fields(line)
		grid = append(grid, words)
		return nil
	})
	sum := 0
	for c := range grid[0] {
		var algFunc func(a, b int) int
		if grid[len(grid)-1][c] == "+" {
			algFunc = add
		} else {
			algFunc = mult
		}
		total, _ := strconv.Atoi(grid[0][c])
		for r := 1; r < len(grid)-1; r++ {
			num, _ := strconv.Atoi(grid[r][c])
			total = algFunc(total, num)
		}
		sum += total
	}
	return sum
}

func main() {
	result := cephMath("../input.txt")
	fmt.Print("RESULT ---\n")
	fmt.Printf("RESULT: %d\n", result)
	fmt.Print("RESULT ---\n")
}

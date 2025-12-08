package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
)

type AlgFunc func(a, b int) int

var add AlgFunc = func(a, b int) int {
	return a + b
}

var mult AlgFunc = func(a, b int) int {
	return a * b
}

func reduce(nums []int, af AlgFunc) int {
	fmt.Printf("reducing %d\n", nums)
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		total = af(total, nums[i])
	}
	fmt.Printf("total %d\n", total)
	return total
}

func cephMath(filePath string) int {
	rows, _ := utils.GetAllInputLines(filePath)
	sum := 0
	algRow := rows[len(rows)-1]
	colNums := []int{}
	var algFunc func(a, b int) int
	for c, char := range algRow {
		if char == '+' {
			algFunc = add
			println("choosing add")
		} else if char == '*' {
			algFunc = mult
			println("choosing mult")
		}
		numStr := ""
		allSpaces := true
		for r := range len(rows) - 1 {
			targetChar := rows[r][c]
			if targetChar != ' ' && targetChar != '\n' {
				numStr += string(targetChar)
				allSpaces = false
			}
		}
		if allSpaces {
			// we've reached the last column in the algFunc
			sum += reduce(colNums, algFunc)
			colNums = []int{}
		} else {
			// otherwise, append the current num
			num, _ := strconv.Atoi(numStr)
			fmt.Printf("appending %d\n", num)
			colNums = append(colNums, num)
			if c == len(algRow)-1 {
				// we've reached the last col in the row
				sum += reduce(colNums, algFunc)
			}
		}
	}
	return sum
}

func main() {
	result := cephMath("../input.txt")
	fmt.Print("RESULT ---\n")
	fmt.Printf("RESULT: %d\n", result)
	fmt.Print("RESULT ---\n")
}

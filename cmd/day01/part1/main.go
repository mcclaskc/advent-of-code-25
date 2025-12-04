package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
)

const max = 99
const numPositions = max + 1

func right(p int, d int) int {
	return (p + d) % numPositions
}

func left(p int, d int) int {
	return right(p, numPositions-(d%numPositions))
}

func main() {
	position := 50
	zeroCount := 0
	utils.ProcessInput("cmd/day01/input.txt", func(line string) error {
		delta, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' { // left
			position = left(position, delta)
		} else {
			position = right(position, delta)
		}
		if position == 0 {
			zeroCount++
		}
		return nil
	})
	fmt.Println(zeroCount)
}

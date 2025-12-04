package main

import (
	"advent-of-code-25/internal/utils"
	"container/ring"
	"fmt"
	"strconv"
)

var zeroCount = 0

func travel(el *ring.Ring, d int, move func(*ring.Ring) *ring.Ring) *ring.Ring {
	for i := 0; i < d; i++ {
		el = move(el)
		if el.Value == 0 {
			zeroCount++
		}
	}
	return el
}

func main() {
	const numPositions = 100
	r := ring.New(numPositions)

	var position *ring.Ring

	for i := 0; i < r.Len(); i++ {
		if i == 50 {
			position = r
		}
		r.Value = i
		r = r.Next()
	}

	utils.ProcessInput("cmd/day01/input.txt", func(line string) error {
		delta, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' { // left
			position = travel(position, delta, (*ring.Ring).Prev)
		} else {
			position = travel(position, delta, (*ring.Ring).Next)
		}
		return nil
	})
	fmt.Println(zeroCount)
}

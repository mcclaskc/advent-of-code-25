package main

import (
	"reflect"
	"strings"
	"testing"
)

const testInputStr = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

type Grid [][]string

func makeGrid(str string) Grid {
	rows := strings.Split(str, "\n")

	grid := Grid{}

	for i := range rows {
		grid = append(grid, strings.Split(rows[i], ""))
	}
	return grid
}

func TestCountTp(t *testing.T) {
	tests := []struct {
		input    Grid
		expected int
	}{
		{makeGrid(`@@@
@@@
@@@`), 4},
		{makeGrid(testInputStr), 13},
	}

	for _, test := range tests {
		t.Run("Grid", func(t *testing.T) {
			actual := countTp(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("findInvalidIdsInRange(%s) = %v but expected %v", test.input, actual, test.expected)
			}
		})
	}
}

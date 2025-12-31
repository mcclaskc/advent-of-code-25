package main

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		joltages Joltages
		buttons  Buttons
		expected Matrix
	}{
		{
			Joltages{5},
			Buttons{Button{0}},
			Matrix{
				[]float64{1, 5},
			},
		},
		{
			Joltages{5, 7, 9},
			Buttons{
				Button{0},
				Button{0, 1},
				Button{0, 1, 2},
			},
			Matrix{
				[]float64{1, 1, 1, 5},
				[]float64{0, 1, 1, 7},
				[]float64{0, 0, 1, 9},
			},
		},
	}

	for _, test := range tests {
		t.Run("JunctionBoxes", func(t *testing.T) {
			actual := generateMatrix(test.joltages, test.buttons)
			if !reflect.DeepEqual(actual, test.expected) {
				printMatrix(actual)
				printMatrix(test.expected)
				t.Errorf("expected %v, actual %v", test.expected, actual)
			}
		})
	}
}

func TestSolveGaussian(t *testing.T) {
	matrix := Matrix{
		[]float64{1, 1, 1, 5},
		[]float64{0, 1, 1, 7},
		[]float64{0, 0, 1, 9},
	}
	expected := Matrix{}
	t.Run("gaussian", func(t *testing.T) {
		_, actual := solveGaussian(matrix)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})
}

func TestCountAllRows(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"../test_input.txt", 33},
	}

	for _, test := range tests {
		t.Run("JunctionBoxes", func(t *testing.T) {
			actual := countAllRows(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestParseLightsToBinary(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"...#.", 8},
		{".##.", 6},
		{".###.#", 46},
	}

	for _, test := range tests {
		t.Run("parselights", func(t *testing.T) {
			actual := parseLightsToBinary([]rune(test.input))
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

func TestCountAllRows(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"../test_input.txt", 7},
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

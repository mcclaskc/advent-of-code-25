package main

import (
	"reflect"
	"testing"
)

func TestCountAllPaths(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"./test_input.txt", 2},
	}

	for _, test := range tests {
		t.Run("JunctionBoxes", func(t *testing.T) {
			actual := countAllPaths(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

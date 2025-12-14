package main

import (
	"reflect"
	"testing"
)

func TestJunctionBoxes(t *testing.T) {
	tests := []struct {
		input    string
		numPairs int
		expected int
	}{
		{"../test_input.txt", 10, 40},
	}

	for _, test := range tests {
		t.Run("JunctionBoxes", func(t *testing.T) {
			actual := junctionBoxes(test.input, test.numPairs)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

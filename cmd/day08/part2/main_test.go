package main

import (
	"reflect"
	"testing"
)

func TestJunctionBoxes(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"../test_input.txt", 25272},
	}

	for _, test := range tests {
		t.Run("JunctionBoxes", func(t *testing.T) {
			actual := junctionBoxes(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %f, actual %f", test.expected, actual)
			}
		})
	}
}

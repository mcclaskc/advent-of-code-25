package main

import (
	"reflect"
	"testing"
)

func TestSplitCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"../test_input.txt", 50},
	}

	for _, test := range tests {
		t.Run("CephMath", func(t *testing.T) {
			actual := biggestArea(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

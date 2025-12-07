package main

import (
	"reflect"
	"testing"
)

func TestCountTp(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"../test_input.txt", 3},
	}

	for _, test := range tests {
		t.Run("Fresh", func(t *testing.T) {
			actual := countFresh(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("countFresh(%s) = %v but expected %v", test.input, actual, test.expected)
			}
		})
	}
}

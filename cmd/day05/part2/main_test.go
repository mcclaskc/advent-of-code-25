package main

import (
	"reflect"
	"testing"
)

func TestCountFresh(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"../test_input.txt", 14},
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

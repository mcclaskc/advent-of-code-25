package main

import (
	"reflect"
	"testing"
)

func TestMaximizeJolts(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := maximizeJolts(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("findInvalidIdsInRange(%s) = %v but expected %v", test.input, actual, test.expected)
			}
		})
	}
}

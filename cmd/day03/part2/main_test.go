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
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := maximizeJolts(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("maximizeJolts(%s) = %v but expected %v", test.input, actual, test.expected)
			}
		})
	}
}

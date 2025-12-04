package main

import (
	"reflect"
	"testing"
)

func TestFindInvalidIdsInRange(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"11-22", []int{11, 22}},
		{"95-115", []int{99}},
		{"998-1012", []int{1010}},
		{"1188511880-1188511890", []int{1188511885}},
		{"222220-222224", []int{222222}},
		{"1698522-1698528", []int{}},
		{"446443-446449", []int{446446}},
		{"38593856-38593862", []int{38593859}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := findInvalidIdsInRange(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("findInvalidIdsInRange(%s) = %v but expected %v", test.input, actual, test.expected)
			}
		})
	}
}

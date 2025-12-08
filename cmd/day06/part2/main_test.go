package main

import (
	"reflect"
	"testing"
)

func TestCephMath(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"../test_input.txt", 3263827},
	}

	for _, test := range tests {
		t.Run("CephMath", func(t *testing.T) {
			actual := cephMath(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, actual %d", test.expected, actual)
			}
		})
	}
}

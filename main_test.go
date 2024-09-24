package main

import (
	"reflect"
	"testing"
)

func TestCharacterOccurrences(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[rune]int
	}{
		{
			name:     "Test 1",
			input:    "Hello",
			expected: map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1},
		},
		{
			name:     "Test 2",
			input:    "Hello, 世界",
			expected: map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1, ',': 1, ' ': 1, '世': 1, '界': 1},
		},
		{
			name:     "Test 3",
			input:    "",
			expected: map[rune]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := characterOccurrences(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

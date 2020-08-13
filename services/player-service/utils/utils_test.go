package utils

import (
	"testing"
)

func TestParseQueryToInt(t *testing.T) {
	tests := []struct {
		name           string
		queries        []string
		expectedLength int
	}{
		{
			name:           "First test",
			queries:        []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
			expectedLength: 10,
		},
		{
			name:           "Second test",
			queries:        []string{"1", "2", "3", "4", "5"},
			expectedLength: 5,
		},
		{
			name:           "Third test",
			queries:        []string{"1", "2", "3", "A"},
			expectedLength: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			intArr, _ := ParseQueryToInt(test.queries...)

			if len(intArr) != test.expectedLength {
				t.Errorf("utils.ParseQueryToInt. \nError: length = %v+ | expected = %v+", len(intArr), test.expectedLength)
			}
		})
	}

}

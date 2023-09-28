package dotenv

import (
	"testing"
)

func TestParseDotenvFile(t *testing.T) {
	tests := []struct {
		input    string
		expected []Variable
	}{
		{
			input: "VAR1=value1\nVAR2=value2",
			expected: []Variable{
				{Key: "VAR1", Value: "value1", Comment: ""},
				{Key: "VAR2", Value: "value2", Comment: ""},
			},
		},
		{
			input: "# Comment\nVAR=value",
			expected: []Variable{
				{Key: "VAR", Value: "value", Comment: "Comment"},
			},
		},
		{
			input: "VAR=value # Comment",
			expected: []Variable{
				{Key: "VAR", Value: "value", Comment: "Comment"},
			},
		},
		{
			input: "VAR=\"value with spaces\" # Comment",
			expected: []Variable{
				{Key: "VAR", Value: "value with spaces", Comment: "Comment"},
			},
		},
		{
			input: "VAR=\"value#with#hashes\"",
			expected: []Variable{
				{Key: "VAR", Value: "value#with#hashes", Comment: ""},
			},
		},
		{
			input: "VAR=value # Comment\nVAR2=value2",
			expected: []Variable{
				{Key: "VAR", Value: "value", Comment: "Comment"},
				{Key: "VAR2", Value: "value2", Comment: ""},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ParseDotenvFile(test.input)
			if len(result) != len(test.expected) {
				t.Errorf("Expected %d variables, but got %d", len(test.expected), len(result))
			}

			for i, variable := range result {
				if variable.Key != test.expected[i].Key {
					t.Errorf("Variable %d: Expected key %s, but got %s", i+1, test.expected[i].Key, variable.Key)
				}
				if variable.Value != test.expected[i].Value {
					t.Errorf("Variable %d: Expected value %s, but got %s", i+1, test.expected[i].Value, variable.Value)
				}
				if variable.Comment != test.expected[i].Comment {
					t.Errorf("Variable %d: Expected comment %s, but got %s", i+1, test.expected[i].Comment, variable.Comment)
				}
			}
		})
	}
}

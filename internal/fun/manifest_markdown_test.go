package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManifestMarkdownErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "no arguments",
			input:    []any{},
			expected: "markdown must be provided",
		},
		{
			name:     "too many arguments",
			input:    []any{"markdown", "extra"},
			expected: "markdown must be provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ManifestMarkdown().Func(tt.input)

			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}

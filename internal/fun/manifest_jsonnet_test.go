package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManifestJsonnetErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "no arguments",
			input:    []any{},
			expected: "jsonnet must be provided",
		},
		{
			name:     "too many arguments",
			input:    []any{"jsonnet", "extra"},
			expected: "jsonnet must be provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ManifestJsonnet().Func(tt.input)

			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}

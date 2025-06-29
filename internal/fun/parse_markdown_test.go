package fun

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		expected string
	}{
		// Basic cases
		{
			name:     "empty string",
			markdown: "",
			expected: `["Document"]`,
		},
		{
			name:     "simple text",
			markdown: "Hello, World!",
			expected: `["Document",["Paragraph","Hello, World","!"]]`,
		},
		{
			name:     "inline code",
			markdown: "This is `inline code` text",
			expected: `["Document",["Paragraph","This is ",["CodeSpan","inline code"]," text"]]`,
		},
		{
			name:     "paragraph with formatting",
			markdown: "This is **bold** and *italic* text.",
			expected: `["Document",["Paragraph","This is ",["Emphasis","bold"]," and ",["Emphasis","italic"]," text."]]`,
		},
		{
			name:     "link",
			markdown: "[Google](https://google.com)",
			expected: `["Document",["Paragraph",["Link","Google","https://google.com"]]]`,
		},
		{
			name:     "image",
			markdown: "![alt text](image.jpg)",
			expected: `["Document",["Paragraph",["Image","image.jpg","alt text"]]]`,
		},
		{
			name:     "horizontal rule",
			markdown: "---",
			expected: `["Document",["ThematicBreak"]]`,
		},
		{
			name:     "blockquote",
			markdown: "> This is a blockquote\n> with multiple lines",
			expected: `["Document",["Blockquote",["Paragraph","This is a blockquote","with multiple lines"]]]`,
		},

		// Headings (progressive levels)
		{
			name:     "heading",
			markdown: "# My Heading",
			expected: `["Document",["Heading1","My Heading"]]`,
		},
		{
			name:     "heading2",
			markdown: "## My Heading",
			expected: `["Document",["Heading2","My Heading"]]`,
		},
		{
			name:     "heading3",
			markdown: "### Heading 3",
			expected: `["Document",["Heading3","Heading 3"]]`,
		},
		{
			name:     "heading4",
			markdown: "#### Heading 4",
			expected: `["Document",["Heading4","Heading 4"]]`,
		},
		{
			name:     "heading5",
			markdown: "##### Heading 5",
			expected: `["Document",["Heading5","Heading 5"]]`,
		},
		{
			name:     "heading6",
			markdown: "###### Heading 6",
			expected: `["Document",["Heading6","Heading 6"]]`,
		},

		// Lists
		{
			name:     "list",
			markdown: "- Item 1\n- Item 2\n- Item 3",
			expected: `["Document",["List",["ListItem",["TextBlock","Item 1"]],["ListItem",["TextBlock","Item 2"]],["ListItem",["TextBlock","Item 3"]]]]`,
		},
		{
			name:     "numbered list",
			markdown: "1. First item\n2. Second item\n3. Third item",
			expected: `["Document",["List",["ListItem",["TextBlock","First item"]],["ListItem",["TextBlock","Second item"]],["ListItem",["TextBlock","Third item"]]]]`,
		},
		{
			name:     "nested list",
			markdown: "- Item 1\n  - Nested item 1\n  - Nested item 2\n- Item 2",
			expected: `["Document",["List",["ListItem",["TextBlock","Item 1"],["List",["ListItem",["TextBlock","Nested item 1"]],["ListItem",["TextBlock","Nested item 2"]]]],["ListItem",["TextBlock","Item 2"]]]]`,
		},

		// Code blocks
		{
			name:     "indented code block",
			markdown: "    func main() {\n        fmt.Println(\"Hello\")\n    }",
			expected: `["Document",["CodeBlock","func main() {\n    fmt.Println(\"Hello\")\n}\n"]]`,
		},
		{
			name:     "code block",
			markdown: "```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```",
			expected: `["Document",["FencedCodeBlock","go","func main() {\n    fmt.Println(\"Hello\")\n}\n"]]`,
		},

		// HTML and special content
		{
			name:     "html block",
			markdown: "<div>HTML content</div>",
			expected: `["Document",["HTMLBlock"]]`,
		},
		{
			name:     "unicode content",
			markdown: "# Unicode Test\n\nHello 世界! 🌍\n\nThis has **émojis** and `código`.",
			expected: `["Document",["Heading1","Unicode Test"],["Paragraph","Hello 世界","! 🌍"],["Paragraph","This has ",["Emphasis","émojis"]," and ",["CodeSpan","código"],"."]]`,
		},

		// Complex combined case
		{
			name:     "complex markdown",
			markdown: "# Title\n\nThis is a **paragraph** with [links](http://example.com) and `code`.\n\n- List item 1\n- List item 2\n\n```python\nprint('Hello')\n```",
			expected: `["Document",["Heading1","Title"],["Paragraph","This is a ",["Emphasis","paragraph"]," with ",["Link","links","http://example.com"]," and ",["CodeSpan","code"],"."],["List",["ListItem",["TextBlock","List item 1"]],["ListItem",["TextBlock","List item 2"]]],["FencedCodeBlock","python","print('Hello')\n"]]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMarkdown().Func([]any{tt.markdown})

			require.NoError(t, err)
			require.NotNil(t, result)

			b, err := json.Marshal(result)
			require.NoError(t, err)

			resultStr := string(b)
			assert.NotEmpty(t, resultStr)

			var jsonData interface{}
			err = json.Unmarshal([]byte(resultStr), &jsonData)
			assert.NoError(t, err, "Result should be valid JSON")

			assert.Equal(t, tt.expected, resultStr, "Result should match expected JSON")
		})
	}
}

func TestParseMarkdownErrors(t *testing.T) {
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
		{
			name:     "non-string argument",
			input:    []any{123},
			expected: "markdown must be a string",
		},
		{
			name:     "nil argument",
			input:    []any{nil},
			expected: "markdown must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMarkdown().Func(tt.input)

			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}

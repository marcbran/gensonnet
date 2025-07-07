package markdown

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseString(t *testing.T) {
	tests := markdownTests()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseString(tt.markdown)

			require.NotNil(t, result)

			b, err := json.Marshal(result)
			require.NoError(t, err)

			resultStr := string(b)
			assert.NotEmpty(t, resultStr)

			var jsonData interface{}
			err = json.Unmarshal([]byte(resultStr), &jsonData)
			assert.NoError(t, err, "Result should be valid JSON")

			assert.Equal(t, tt.element, resultStr, "Result should match element JSON")
		})
	}
}

func TestManifestAny(t *testing.T) {
	tests := markdownTests()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := ParseString(tt.markdown)
			resultStr, err := ManifestAny(node)
			require.NoError(t, err)

			canonical := tt.canonical
			if canonical == "" {
				canonical = tt.markdown
			}

			assert.Equal(t, canonical, strings.TrimSuffix(resultStr, "\n"), "Result should match element JSON")
		})
	}
}

func markdownTests() []struct {
	name      string
	markdown  string
	element   string
	canonical string
} {
	return []struct {
		name      string
		markdown  string
		element   string
		canonical string
	}{
		// Basic cases
		{
			name:     "empty string",
			markdown: "",
			element:  `["Document"]`,
		},
		{
			name:     "string",
			markdown: "string",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},"string"]]`,
		},
		{
			name:     "simple text",
			markdown: "Hello, World!",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},"Hello, World","!"]]`,
		},
		{
			name:     "soft line break",
			markdown: "Hello, World!\nWelcome to my website!",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},"Hello, World",["Text",{"softLineBreak":true}],"Welcome to my website","!"]]`,
		},
		{
			name:      "hard line break",
			markdown:  "Hello, World!  \nWelcome to my website!",
			element:   `["Document",["Paragraph",{"blankPreviousLines":true},"Hello, World",["Text",{"hardLineBreak":true}],"Welcome to my website","!"]]`,
			canonical: "Hello, World!\\\nWelcome to my website!",
		},
		{
			name:     "hard line break canonical",
			markdown: "Hello, World!\\\nWelcome to my website!",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},"Hello, World",["Text",{"hardLineBreak":true}],"Welcome to my website","!"]]`,
		},
		{
			name:     "inline code",
			markdown: "This is `inline code` text",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},"This is ",["CodeSpan","inline code"]," text"]]`,
		},
		{
			name:     "paragraph with formatting",
			markdown: "This is **bold** and *italic* text.",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},"This is ",["Emphasis",{"level":2},"bold"]," and ",["Emphasis",{"level":1},"italic"]," text."]]`,
		},
		{
			name:     "link",
			markdown: "[Google](https://google.com)",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},["Link",{"destination":"https://google.com"},"Google"]]]`,
		},
		{
			name:     "image",
			markdown: "![alt text](image.jpg)",
			element:  `["Document",["Paragraph",{"blankPreviousLines":true},["Image",{"destination":"image.jpg"},"alt text"]]]`,
		},
		{
			name:     "horizontal rule",
			markdown: "---",
			element:  `["Document",["ThematicBreak",{"blankPreviousLines":true}]]`,
		},
		{
			name:     "blockquote",
			markdown: "> This is a blockquote\n> with multiple lines",
			element:  `["Document",["Blockquote",{"blankPreviousLines":true},["Paragraph",{"blankPreviousLines":true},["Text",{"softLineBreak":true}],"with multiple lines"]]]`,
		},

		// Headings (progressive levels)
		{
			name:     "heading",
			markdown: "# My Heading",
			element:  `["Document",["Heading",{"level":1},"My Heading"]]`,
		},
		{
			name:     "heading2",
			markdown: "## My Heading",
			element:  `["Document",["Heading",{"level":2},"My Heading"]]`,
		},
		{
			name:     "heading3",
			markdown: "### Heading 3",
			element:  `["Document",["Heading",{"level":3},"Heading 3"]]`,
		},
		{
			name:     "heading4",
			markdown: "#### Heading 4",
			element:  `["Document",["Heading",{"level":4},"Heading 4"]]`,
		},
		{
			name:     "heading5",
			markdown: "##### Heading 5",
			element:  `["Document",["Heading",{"level":5},"Heading 5"]]`,
		},
		{
			name:     "heading6",
			markdown: "###### Heading 6",
			element:  `["Document",["Heading",{"level":6},"Heading 6"]]`,
		},

		// Lists
		{
			name:     "list",
			markdown: "- Item 1\n- Item 2\n- Item 3",
			element:  `["Document",["List",{"blankPreviousLines":true,"marker":"-","start":0},["ListItem",{"blankPreviousLines":true},["TextBlock","Item 1"]],["ListItem",["TextBlock","Item 2"]],["ListItem",["TextBlock","Item 3"]]]]`,
		},
		{
			name:     "numbered list",
			markdown: "1. First item\n2. Second item\n3. Third item",
			element:  `["Document",["List",{"blankPreviousLines":true,"marker":".","start":1},["ListItem",{"blankPreviousLines":true},["TextBlock","First item"]],["ListItem",["TextBlock","Second item"]],["ListItem",["TextBlock","Third item"]]]]`,
		},
		{
			name:     "nested list",
			markdown: "- Item 1\n  - Nested item 1\n  - Nested item 2\n- Item 2",
			element:  `["Document",["List",{"blankPreviousLines":true,"marker":"-","start":0},["ListItem",{"blankPreviousLines":true},["TextBlock","Item 1"],["List",{"marker":"-","start":0},["ListItem",["TextBlock","Nested item 1"]],["ListItem",["TextBlock","Nested item 2"]]]],["ListItem",["TextBlock","Item 2"]]]]`,
		},

		// Code Blocks
		{
			name:     "code block",
			markdown: "    func main() {\n        fmt.Println(\"Hello\")\n    }",
			element:  `["Document",["CodeBlock",{"blankPreviousLines":true},"func main() {\n    fmt.Println(\"Hello\")\n}"]]`,
		},
		{
			name:     "fenced code block",
			markdown: "```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```",
			element:  `["Document",["FencedCodeBlock",{"blankPreviousLines":true,"language":"go"},"func main() {\n    fmt.Println(\"Hello\")\n}"]]`,
		},

		// HTML and special content
		{
			name:     "html block",
			markdown: "<div>HTML content</div>",
			element:  `["Document",["HTMLBlock",{"blankPreviousLines":true},"\u003cdiv\u003eHTML content\u003c/div\u003e"]]`,
		},
		{
			name:     "unicode content",
			markdown: "# Unicode Test\n\nHello 世界! 🌍\n\nThis has **émojis** and `código`.",
			element:  `["Document",["Heading",{"level":1},"Unicode Test"],["Paragraph",{"blankPreviousLines":true},"Hello 世界","! 🌍"],["Paragraph",{"blankPreviousLines":true},"This has ",["Emphasis",{"level":2},"émojis"]," and ",["CodeSpan","código"],"."]]`,
		},

		// Complex combined case
		{
			name:     "complex markdown",
			markdown: "# Title\n\nThis is a **paragraph** with [links](http://example.com) and `code`.\n\n- List item 1\n- List item 2\n\n```python\nprint('Hello')\n```",
			element:  `["Document",["Heading",{"level":1},"Title"],["Paragraph",{"blankPreviousLines":true},"This is a ",["Emphasis",{"level":2},"paragraph"]," with ",["Link",{"destination":"http://example.com"},"links"]," and ",["CodeSpan","code"],"."],["List",{"blankPreviousLines":true,"marker":"-","start":0},["ListItem",{"blankPreviousLines":true},["TextBlock","List item 1"]],["ListItem",["TextBlock","List item 2"]]],["FencedCodeBlock",{"blankPreviousLines":true,"language":"python"},"print('Hello')"]]`,
		},
	}
}

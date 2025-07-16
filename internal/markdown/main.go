package markdown

import (
	"bytes"
	"strings"

	markdown "github.com/teekennedy/goldmark-markdown"
	"github.com/yuin/goldmark"
	mdast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func ParseString(val string) any {
	node := goldmark.DefaultParser().Parse(text.NewReader([]byte(val)))
	out := Parse(node, []byte(val))
	return out
}

func Parse(node mdast.Node, source []byte) any {
	attributes := make(map[string]any)
	var value string
	switch n := node.(type) {
	case *mdast.Blockquote:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
	case *mdast.CodeBlock:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
		value = linesValue(node, source)
	case *mdast.Document:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
	case *mdast.Emphasis:
		attributes["level"] = float64(n.Level)
	case *mdast.FencedCodeBlock:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
		attributes["language"] = string(n.Language(source))
		value = linesValue(node, source)
	case *mdast.Heading:
		attributes["level"] = float64(n.Level)
	case *mdast.HTMLBlock:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
		value = linesValue(node, source)
	case *mdast.Image:
		attributes["destination"] = string(n.Destination)
	case *mdast.Link:
		attributes["destination"] = string(n.Destination)
	case *mdast.List:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
		attributes["marker"] = string(n.Marker)
		attributes["start"] = float64(n.Start)
	case *mdast.ListItem:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
	case *mdast.Paragraph:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
	case *mdast.ThematicBreak:
		if n.HasBlankPreviousLines() {
			attributes["blankPreviousLines"] = n.HasBlankPreviousLines()
		}
	case *mdast.Text:
		if n.SoftLineBreak() {
			attributes["softLineBreak"] = n.SoftLineBreak()
		}
		if n.HardLineBreak() {
			attributes["hardLineBreak"] = n.HardLineBreak()
		}
		value = string(n.Value(source))
		if len(attributes) == 0 {
			return value
		}
	default:
	}
	res := []any{tag(node)}
	if len(attributes) > 0 {
		res = append(res, attributes)
	}
	if value != "" {
		res = append(res, value)
	}
	res = append(res, parseChildren(node, source)...)
	return res
}

func tag(node mdast.Node) string {
	return node.Kind().String()
}

func linesValue(node mdast.Node, source []byte) string {
	return strings.TrimSuffix(string(node.Lines().Value(source)), "\n")
}

func parseChildren(node mdast.Node, source []byte) []any {
	var res []any
	child := node.FirstChild()
	for child != nil {
		converted := Parse(child, source)
		res = append(res, converted)
		child = child.NextSibling()
	}
	return res
}

func ManifestAny(elem any) (string, error) {
	node, source, err := Manifest(elem)
	if err != nil {
		return "", err
	}
	gm := goldmark.New(goldmark.WithRenderer(markdown.NewRenderer()))
	builder := strings.Builder{}
	err = gm.Renderer().Render(&builder, source, node)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func Manifest(elem any) (mdast.Node, []byte, error) {
	buffer := bytes.Buffer{}
	node, err := manifestRec(elem, &buffer)
	if err != nil {
		return nil, nil, err
	}
	return node, buffer.Bytes(), nil
}

func manifestRec(elem any, buffer *bytes.Buffer) (mdast.Node, error) {
	var err error
	switch v := elem.(type) {
	case []any:
		if len(v) == 0 {
			return nil, nil
		}
		tag, attributes, children := elemParts(v)
		node := newNode(tag)
		switch n := node.(type) {
		case *mdast.Blockquote:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
		case *mdast.CodeBlock:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
			if len(children) > 0 {
				segments, err := addSegments(buffer, children[0].(string))
				if err != nil {
					return nil, err
				}
				n.SetLines(&segments)
			}
			children = nil
		case *mdast.Document:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
		case *mdast.Emphasis:
			if level, ok := attributes["level"].(float64); ok {
				n.Level = int(level)
			}
		case *mdast.FencedCodeBlock:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
			if language, ok := attributes["language"].(string); ok {
				segment, err := addSegment(buffer, language)
				if err != nil {
					return nil, err
				}
				n.Info = mdast.NewTextSegment(segment)
			}
			if len(children) > 0 {
				segments, err := addSegments(buffer, children[0].(string))
				if err != nil {
					return nil, err
				}
				n.SetLines(&segments)
			}
			children = nil
		case *mdast.Heading:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
			if level, ok := attributes["level"].(float64); ok {
				n.Level = int(level)
			}
		case *mdast.HTMLBlock:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
			if len(children) > 0 {
				segments, err := addSegments(buffer, children[0].(string))
				if err != nil {
					return nil, err
				}
				n.SetLines(&segments)
			}
			children = nil
		case *mdast.Image:
			if destination, ok := attributes["destination"].(string); ok {
				n.Destination = []byte(destination)
			}
		case *mdast.Link:
			if destination, ok := attributes["destination"].(string); ok {
				n.Destination = []byte(destination)
			}
		case *mdast.List:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
			if marker, ok := attributes["marker"].(string); ok {
				n.Marker = marker[0]
			}
			if start, ok := attributes["start"].(float64); ok {
				n.Start = int(start)
			}
		case *mdast.ListItem:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
		case *mdast.Paragraph:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
		case *mdast.ThematicBreak:
			if blankPreviousLines, ok := attributes["blankPreviousLines"].(bool); ok {
				n.SetBlankPreviousLines(blankPreviousLines)
			}
		case *mdast.Text:
			if softLineBreak, ok := attributes["softLineBreak"].(bool); ok {
				n.SetSoftLineBreak(softLineBreak)
			}
			if hardLineBreak, ok := attributes["hardLineBreak"].(bool); ok {
				n.SetHardLineBreak(hardLineBreak)
			}
			if len(children) > 0 {
				n.Segment, err = addSegment(buffer, children[0].(string))
				if err != nil {
					return nil, err
				}
			}
			children = nil
		default:
		}
		err := manifestChildren(node, buffer, children)
		if err != nil {
			return nil, err
		}
		return node, nil
	case string:
		segment, err := addSegment(buffer, v)
		if err != nil {
			return nil, err
		}
		return mdast.NewTextSegment(segment), nil
	}
	return nil, nil
}

func manifestChildren(node mdast.Node, buffer *bytes.Buffer, children []any) error {
	for _, child := range children {
		childNode, err := manifestRec(child, buffer)
		if err != nil {
			return err
		}
		node.AppendChild(node, childNode)
	}
	return nil
}

func elemParts(val []any) (string, map[string]any, []any) {
	if len(val) == 0 {
		return "", nil, nil
	}
	tag := val[0].(string)
	if len(val) == 1 {
		return tag, nil, nil
	}
	childrenStart := 1
	var attributes map[string]any
	if a, ok := val[1].(map[string]any); ok {
		childrenStart = 2
		attributes = a
	}
	return tag, attributes, val[childrenStart:]
}

func newNode(tag string) mdast.Node {
	switch tag {
	case "Blockquote":
		return mdast.NewBlockquote()
	case "CodeBlock":
		return mdast.NewCodeBlock()
	case "CodeSpan":
		return mdast.NewCodeSpan()
	case "Document":
		return mdast.NewDocument()
	case "Emphasis":
		return &mdast.Emphasis{}
	case "FencedCodeBlock":
		return &mdast.FencedCodeBlock{}
	case "Heading":
		return &mdast.Heading{}
	case "HTMLBlock":
		return &mdast.HTMLBlock{}
	case "Image":
		return &mdast.Image{}
	case "Link":
		return mdast.NewLink()
	case "List":
		return &mdast.List{}
	case "ListItem":
		return &mdast.ListItem{}
	case "Paragraph":
		return mdast.NewParagraph()
	case "RawHTML":
		return mdast.NewRawHTML()
	case "Text":
		return mdast.NewText()
	case "TextBlock":
		return mdast.NewTextBlock()
	case "ThematicBreak":
		return mdast.NewThematicBreak()
	}
	return nil
}

func addSegments(buffer *bytes.Buffer, val string) (text.Segments, error) {
	lines := strings.Split(val, "\n")
	var segments text.Segments
	for _, line := range lines {
		segment, err := addSegment(buffer, line)
		if err != nil {
			return text.Segments{}, err
		}
		segments.Append(segment)
	}
	return segments, nil
}

func addSegment(buffer *bytes.Buffer, val string) (text.Segment, error) {
	start := buffer.Len()
	n, err := buffer.WriteString(val)
	if err != nil {
		return text.Segment{}, err
	}
	return text.NewSegment(start, start+n), nil
}

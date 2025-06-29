package fun

import (
	"fmt"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/yuin/goldmark"
	mdast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func ParseMarkdown() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "parseMarkdown",
		Params: ast.Identifiers{"markdown"},
		Func: func(input []any) (any, error) {
			if len(input) != 1 {
				return nil, fmt.Errorf("markdown must be provided")
			}
			markdown, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("markdown must be a string")
			}
			node := goldmark.DefaultParser().Parse(text.NewReader([]byte(markdown)))
			out := convert(node, []byte(markdown))
			return out, nil
		},
	}
}

func convert(node mdast.Node, source []byte) any {
	switch node.Kind() {
	case mdast.KindLink:
		link := node.(*mdast.Link)
		return []any{
			tag(link),
			convert(link.FirstChild(), source),
			string(link.Destination),
		}
	case mdast.KindImage:
		image := node.(*mdast.Image)
		return []any{
			tag(image),
			string(image.Destination),
			convert(image.FirstChild(), source),
		}
	case mdast.KindText:
		return string(node.(*mdast.Text).Value(source))
	case mdast.KindCodeBlock:
		block := node.(*mdast.CodeBlock)
		return []any{
			tag(block),
			string(block.BaseBlock.Lines().Value(source)),
		}
	case mdast.KindFencedCodeBlock:
		block := node.(*mdast.FencedCodeBlock)
		return []any{
			tag(block),
			string(block.Language(source)),
			string(block.BaseBlock.Lines().Value(source)),
		}
	default:
		res := []any{tag(node)}
		child := node.FirstChild()
		for child != nil {
			converted := convert(child, source)
			res = append(res, converted)
			child = child.NextSibling()
		}
		return res
	}
}

func tag(node mdast.Node) string {
	switch node.Kind() {
	case mdast.KindHeading:
		return fmt.Sprintf("Heading%d", node.(*mdast.Heading).Level)
	default:
		return node.Kind().String()
	}
}

package fun

import (
	"fmt"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/marcbran/gensonnet/internal/markdown"
)

func ParseMarkdown() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "parseMarkdown",
		Params: ast.Identifiers{"markdown"},
		Func: func(input []any) (any, error) {
			if len(input) != 1 {
				return nil, fmt.Errorf("markdown must be provided")
			}
			md, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("markdown must be a string")
			}
			out := markdown.ParseString(md)
			return out, nil
		},
	}
}

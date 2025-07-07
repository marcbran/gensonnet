package fun

import (
	"fmt"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	"github.com/marcbran/gensonnet/internal/fun/markdown"
)

func ManifestMarkdown() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "manifestMarkdown",
		Params: ast.Identifiers{"markdown"},
		Func: func(input []any) (any, error) {
			if len(input) != 1 {
				return nil, fmt.Errorf("markdown must be provided")
			}
			out, err := markdown.ManifestAny(input[0])
			if err != nil {
				return nil, err
			}
			return out, nil
		},
	}
}

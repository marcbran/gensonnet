package fun

import (
	"fmt"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
	intjsonnet "github.com/marcbran/gensonnet/internal/jsonnet"
)

func ManifestJsonnet() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "manifestJsonnet",
		Params: ast.Identifiers{"jsonnet"},
		Func: func(input []any) (any, error) {
			if len(input) != 1 {
				return nil, fmt.Errorf("jsonnet must be provided")
			}
			out, err := intjsonnet.Manifest(input[0])
			if err != nil {
				return nil, err
			}
			return out, nil
		},
	}
}

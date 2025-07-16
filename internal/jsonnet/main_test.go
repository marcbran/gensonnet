package jsonnet

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseManifest(t *testing.T) {
	tests := []struct {
		name      string
		jsonnet   string
		canonical string
	}{
		{
			name:    "null/simple",
			jsonnet: "null",
		},
		{
			name:    "true/simple",
			jsonnet: "true",
		},
		{
			name:    "false/simple",
			jsonnet: "false",
		},
		{
			name:    "self/simple",
			jsonnet: "{ a: 'foo', b: self.a }",
		},
		{
			name:    "outer",
			jsonnet: "{ a: 'foo', b: { c: $.a } }",
		},
		{
			name:    "super",
			jsonnet: "{ a: 'foo' } + { b: super.a }",
		},
		{
			name:    "string/single quote",
			jsonnet: "'foobar'",
		},
		{
			name:    "string/format",
			jsonnet: "'foobar %s' % ['baz']",
		},
		{
			name:    "number/single digit",
			jsonnet: "3",
		},
		{
			name:    "number/decimal",
			jsonnet: "3.123",
		},
		{
			name:    "id/simple",
			jsonnet: "local a = 1; a",
		},
		{
			name:    "member/id member",
			jsonnet: "local a = { b: 1 }; a.b",
		},
		{
			name:    "member/id nested member",
			jsonnet: "local a = { b: { c: 1 } }; a.b.c",
		},
		{
			name:    "index/id index",
			jsonnet: "local a = { 'foo-bar': 1 }; a['foo-bar']",
		},
		{
			name:    "index/id nested index",
			jsonnet: "local a = { 'foo-bar': { 'bar-foo': 1 } }; a['foo-bar']['bar-foo']",
		},
		{
			name:    "func/one param",
			jsonnet: "std.map(function(a) a, [1, 2, 3])",
		},
		{
			name:    "func/one default param",
			jsonnet: "std.map(function(a=2) a, [1, 2, 3])",
		},
		{
			name:    "call/no params",
			jsonnet: "local foo() = 1; foo()",
		},
		{
			name:    "call/one params",
			jsonnet: "local a = 1; local foo(a) = a; foo(a)",
		},
		{
			name:    "call/one named params",
			jsonnet: "local a = 1; local foo(a) = a; foo(a=2)",
		},
		{
			name:    "call/two params",
			jsonnet: "local a = 1; local foo(a, b) = a + b; foo(a, 1)",
		},
		{
			name:    "call/two params, one named param",
			jsonnet: "local a = 1; local foo(a, b) = a + b; foo(a, b=1)",
		},
		{
			name:    "object/empty",
			jsonnet: "{}",
		},
		{
			name:    "object/field",
			jsonnet: "{ a: 1 }",
		},
		{
			name:    "object/local",
			jsonnet: "{ local a = 1, b: a }",
		},
		{
			name:    "object/assert",
			jsonnet: "{ local a = 1, assert a == 1 : 'a must be 1', b: a }",
		},
		{
			name:    "objectComp/single",
			jsonnet: "{ [number]: number + 1 for number in ['1', '2', '3'] }",
		},
		{
			name:    "objectComp/double",
			jsonnet: "{ [a + b]: a + b for a in ['1', '2', '3'] for b in ['4', '5', '6'] }",
		},
		{
			name:    "objectComp/for if",
			jsonnet: "{ [a]: a + 1 for a in ['1', '2', '3'] if a == '2' }",
		},
		{
			name:    "array/empty",
			jsonnet: "[]",
		},
		{
			name:    "array/single",
			jsonnet: "['a']",
		},
		{
			name:    "array/local",
			jsonnet: "[local a = 1; a]",
		},
		{
			name:    "arrayComp/single",
			jsonnet: "[number for number in [1, 2, 3]]",
		},
		{
			name:    "arrayComp/double",
			jsonnet: "[a + b for a in [1, 2, 3] for b in [4, 5, 6]]",
		},
		{
			name:    "arrayComp/for if",
			jsonnet: "[a + 1 for a in [1, 2, 3] if a > 2]",
		},
		{
			name:    "if/if",
			jsonnet: "if 1 > 0 then 1",
		},
		{
			name:    "if/if-else",
			jsonnet: "if 1 > 0 then 1 else 0",
		},
		{
			name:    "comment/simple",
			jsonnet: "// Output the number one\n1",
		},
		{
			name:      "comment/in object",
			jsonnet:   "{a: 1,  // b is equal to two\n  b: 2}",
			canonical: "{\n  a: 1,  // b is equal to two\n  b: 2,\n}",
		},
		{
			name:      "comment/in array",
			jsonnet:   "[1, // add two to array\n2]",
			canonical: "[\n  1,  // add two to array\n  2,\n]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, err := Parse(tt.jsonnet)
			require.NoError(t, err)

			resultStr, err := Manifest(node)
			require.NoError(t, err)

			canonical := tt.canonical
			if canonical == "" {
				canonical = tt.jsonnet
			}

			assert.Equal(t, canonical, strings.TrimSuffix(resultStr, "\n"), "Result should match jsonnet")
		})
	}
}

package main_test

import (
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/require"
)

type Statement struct {
	Key   string `parser:"@Ident '='"`
	Value Value  `parser:"@@"`
}

type Value struct {
	String string     `parser:"  @String"`
	Float  float64    `parser:"| @Float"`
	Int    int        `parser:"| @Int"`
	Bool   *BoolValue `parser:"| @('true' | 'false')"`
}

type BoolValue bool

func (b *BoolValue) Capture(values []string) error {
	*b = strings.ToLower(values[0]) == "true"
	return nil
}

func TestField(t *testing.T) {
	parser := participle.MustBuild[Statement]()
	ast, err := parser.ParseString("", "size=10")
	if err != nil {
		panic(err)
	}

	require.Equal(t, "size", ast.Key)
	require.Equal(t, 10, ast.Value.Int)
}

type Block struct {
	Identifier string      `parser:"'define' @Ident"`
	Statements []Statement `parser:"'{' @@* '}'"`
}

func TestBlock(t *testing.T) {
	parser := participle.MustBuild[Block](participle.Unquote("String"))
	ast, err := parser.ParseString("", `
		define  block_1 {
			size   = 10
			length = 20
			flag   = true
			text   = "asdf"
		}
	`)

	bTrue := BoolValue(true)
	expect := &Block{
		Identifier: "block_1",
		Statements: []Statement{
			{Key: "size", Value: Value{Int: 10}},
			{Key: "length", Value: Value{Int: 20}},
			{Key: "flag", Value: Value{Bool: &bTrue}},
			{Key: "text", Value: Value{String: "asdf"}},
		},
	}
	if err != nil {
		panic(err)
	}
	require.NotNil(t, ast)
	require.Equal(t, expect, ast)
}

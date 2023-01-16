package main

import (
	"strings"

	"github.com/alecthomas/participle/v2"
)

func NewParser() Parser {
	// cmdlexer := lexer.Must(lexer.Rules{
	// 	"Root": {
	// 		{`String`, `"`, lexer.Push("String")},
	// 	},
	// 	{Name: "Comment", Pattern: `#[^\n]*`},
	// 	{Name: "String", Pattern: `"(\\"|[^"])*"`},
	// 	{Name: "Number", Pattern: `[-+]?(\d*\.)?\d+`},
	// 	{Name: "Ident", Pattern: `[a-zA-Z_]\w*`},
	// 	{Name: "Punct", Pattern: `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	// 	{Name: "EOL", Pattern: `[\n\r]+`},
	// 	{Name: "whitespace", Pattern: `[ \t]+`},
	// 	{Name: "cmdShow", Pattern: `Show`},
	// })
	build := participle.MustBuild[CmdLine](
		// participle.Lexer(cmdlexer),
		participle.Unquote("String"),
		participle.Union[Cmd](ShowCmd{}, ExitCmd{}),
		//participle.CaseInsensitive("CmdShow", "Object"),
	)
	p := &parser{
		build: build,
	}
	return p
}

type CmdLine struct {
	Cmd Cmd `parser:"@@*"`
	// Flags   []Flag `parser:""`
	// Args []Arg `parser:""`
}

type Cmd interface {
	Name() string
}

// ///////////////////////////
// Exit
type ExitCmd struct {
	Command string `parser:"@('exit')"`
}

func (exit ExitCmd) Name() string { return "EXIT" }

// ///////////////////////////
// Show
type ShowCmd struct {
	Command string `parser:"@('show'|'Show')"`
	Object  string `parser:"(@( 'tables' | 'version' ))?"`
}

func (show ShowCmd) Name() string { return "SHOW" }

// ///////////////////////////
// Others

type Flag struct {
	Key   Key
	Value Value
}

type Arg struct {
}

type Key struct {
	Name *string `parser:" @Ident '='"`
}

type Value struct {
	String *string    `parser:"  @String"`
	Float  *float64   `parser:"| @Float"`
	Int    *int       `parser:"| @Int"`
	Bool   *BoolValue `parser:"| @('true' | 'false')"`
}

type BoolValue bool

func (b *BoolValue) Capture(values []string) error {
	*b = strings.ToLower(values[0]) == "true"
	return nil
}

package main

import (
	"bytes"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Parser interface {
	Parse(line string) (*CmdLine, error)
	Lex(line string) ([]lexer.Token, error)
}

type parser struct {
	build *participle.Parser[CmdLine]
}

func (p *parser) Parse(line string) (*CmdLine, error) {
	cl, err := p.build.ParseString("", line)
	return cl, err
}

func (p *parser) Lex(line string) ([]lexer.Token, error) {
	return p.build.Lex("", bytes.NewBuffer([]byte(line)))
}

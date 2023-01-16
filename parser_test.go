package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseExit(t *testing.T) {
	ps := NewParser()

	// exit
	cmdLine, err := ps.Parse("ExIT")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	require.NotNil(t, cmdLine)
}

func TestParseShow(t *testing.T) {
	var ps = NewParser()
	var cmdLine *CmdLine
	var show ShowCmd
	var err error

	// show
	cmdLine, err = ps.Parse("Show")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	require.NotNil(t, cmdLine)
	require.Equal(t, "SHOW", cmdLine.Cmd.Name())
	show = cmdLine.Cmd.(ShowCmd)
	require.Equal(t, "", show.Object)

	// show tables
	cmdLine, err = ps.Parse("Show tables ")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	require.NotNil(t, cmdLine)
	require.Equal(t, "SHOW", cmdLine.Cmd.Name())
	show = cmdLine.Cmd.(ShowCmd)
	require.Equal(t, "TABLES", strings.ToUpper(show.Object))

	// show tables
	cmdLine, err = ps.Parse("show version ")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	require.NotNil(t, cmdLine)
	require.Equal(t, "SHOW", cmdLine.Cmd.Name())
	show = cmdLine.Cmd.(ShowCmd)
	require.Equal(t, "VERSION", strings.ToUpper(show.Object))
}

func TestLexShow(t *testing.T) {
	var ps = NewParser()

	tokens, err := ps.Lex("Show")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	for i, t := range tokens {
		fmt.Println(i, t.Pos, t.Value, t.Type)
	}
}

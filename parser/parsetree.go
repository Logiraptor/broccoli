package parser

import (
	"github.com/Logiraptor/broccoli/lexer"
)

type Type int

const (
	IdentTree Type = iota
	NumberTree
	SExprTree
)

type Tree interface {
	Type() Type
	Start() lexer.Pos
	End() lexer.Pos
}

type Parser interface {
	Next() (Tree, error)
}

type base struct {
	StartPos, EndPos lexer.Pos
}

func (b base) Start() lexer.Pos {
	return b.StartPos
}

func (b base) End() lexer.Pos {
	return b.EndPos
}

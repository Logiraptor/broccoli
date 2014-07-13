package parser

import (
	"github.com/Logiraptor/broccoli/lexer"
)

func NewIdentifier(name string, start, end lexer.Pos) Tree {
	return Identifier{
		Name: name,
		base: base{
			StartPos: start,
			EndPos:   end,
		},
	}
}

type Identifier struct {
	base
	Name string
}

func (i Identifier) Type() Type {
	return IdentTree
}

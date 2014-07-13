package parser

import (
	"github.com/Logiraptor/broccoli/lexer"
)

func NewSExpression(start, end lexer.Pos, elems ...Tree) Tree {
	return SExpression{
		base: base{
			StartPos: start,
			EndPos:   end,
		},
		Elems: elems,
	}
}

type SExpression struct {
	base
	Elems []Tree
}

func (s SExpression) Type() Type {
	return SExprTree
}

package parser

import (
	"github.com/Logiraptor/broccoli/lexer"
)

func NewNumber(val string, start, end lexer.Pos) Tree {
	return Number{
		Val: val,
		base: base{
			StartPos: start,
			EndPos:   end,
		},
	}
}

type Number struct {
	base
	Val string
}

func (i Number) Type() Type {
	return NumberTree
}

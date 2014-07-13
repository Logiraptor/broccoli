package lexer

type base struct {
	start Pos
	end   Pos
	text  string
	typ   Type
}

func (b base) Start() Pos {
	return b.start
}

func (b base) End() Pos {
	return b.end
}

func (b base) Text() string {
	return b.text
}

func (b base) Type() Type {
	return b.typ
}

const (
	ErrorToken Type = iota
	IdentToken
	OpenToken
	CloseToken
	QuoteToken
	ConsToken
)

package lexer

type Pos struct {
	Line   int
	Column int
	Pos    int
}

type Type int

// Token is a lexed token
type Token interface {
	Start() Pos
	End() Pos
	Type() Type
	Text() string
}

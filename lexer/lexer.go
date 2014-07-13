package lexer

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp/syntax"
	"unicode"
	"unicode/utf8"
)

const eof = -1

var Done = errors.New("done lexing")

type Lexer interface {
	Next() (Token, error)
}

func Lex(source io.Reader) Lexer {
	buf, err := ioutil.ReadAll(source)
	if err != nil {
		panic(err)
	}

	l := &lexer{
		source: buf,
		tokens: make(chan Token),
	}
	go l.run()

	return l
}

type lexer struct {
	source []byte
	tokens chan Token

	lastPos, lastLine, lastCol    int
	pos, line, col                int
	startPos, startLine, startCol int
}

func (l *lexer) next() rune {
	var r rune
	var length int
	if l.pos == len(l.source) {
		r = eof
		length = 0
	} else {
		r, length = utf8.DecodeRune(l.source[l.pos:])
	}
	l.lastPos = l.pos
	l.lastLine = l.line
	l.lastCol = l.col

	l.pos += length
	l.col++
	if r == '\n' {
		l.line++
		l.col = 0
	}
	return r
}

func (l *lexer) backup() {
	l.pos = l.lastPos
	l.line = l.lastLine
	l.col = l.lastCol
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) ignore() {
	l.lastCol = l.col
	l.lastLine = l.line
	l.lastPos = l.pos
	l.startLine = l.line
	l.startPos = l.pos
	l.startCol = l.col
}

func (l *lexer) Next() (Token, error) {
	t, ok := <-l.tokens
	if !ok {
		return nil, Done
	}
	if t.Type() == ErrorToken {
		return nil, errors.New(t.Text())
	}
	return t, nil
}

func (l *lexer) run() {
	state := lexMain
	for state != nil {
		state = state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(typ Type) {
	l.tokens <- base{
		start: Pos{
			Line:   l.startLine,
			Column: l.startCol,
			Pos:    l.startPos,
		},
		end: Pos{
			Line:   l.line,
			Column: l.col,
			Pos:    l.pos,
		},
		text: string(l.source[l.startPos:l.pos]),
		typ:  typ,
	}
	l.lastCol = l.col
	l.lastLine = l.line
	l.lastPos = l.pos
	l.startLine = l.line
	l.startPos = l.pos
	l.startCol = l.col
}

func (l *lexer) errorf(format string, args ...interface{}) {
	l.tokens <- base{
		start: Pos{
			Line:   l.startLine,
			Column: l.startCol,
			Pos:    l.startPos,
		},
		end: Pos{
			Line:   l.line,
			Column: l.col,
			Pos:    l.pos,
		},
		text: fmt.Sprintf("line %d col %d: %s", l.line, l.col, fmt.Sprintf(format, args...)),
		typ:  ErrorToken,
	}
	l.lastCol = l.col
	l.lastLine = l.line
	l.lastPos = l.pos
	l.startLine = l.line
	l.startPos = l.pos
	l.startCol = l.col
}

type stateFn func(l *lexer) stateFn

func lexMain(l *lexer) stateFn {
	switch r := l.next(); {
	case unicode.IsSpace(r):
		l.ignore()
		return lexMain
	case isIdentifier(r):
		return lexIdentifier
	case r == '(':
		l.emit(OpenToken)
		return lexMain
	case r == ')':
		l.emit(CloseToken)
		return lexMain
	case r == '\'':
		l.emit(QuoteToken)
		return lexMain
	case r == '.':
		l.emit(ConsToken)
		return lexMain
	case r == eof:
		return nil
	default:
		l.errorf("invalid rune %d -> %q", r, string(r))
	}
	return nil
}

func isIdentifier(r rune) bool {
	switch {
	case syntax.IsWordChar(r):
		return true
	case r == '?':
		return true
	case r == '+', r == '-', r == '*', r == '/':
		return true
	}
	return false
}

func lexIdentifier(l *lexer) stateFn {
	for {
		switch r := l.peek(); {
		case isIdentifier(r):
			l.next()
		case r == '.':
			l.next()
		default:
			l.emit(IdentToken)
			return lexMain
		}
	}
}

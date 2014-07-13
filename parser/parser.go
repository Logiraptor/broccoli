package parser

import (
	"fmt"
	"io"
	"regexp"

	"github.com/Logiraptor/broccoli/lexer"
)

var numRegexp *regexp.Regexp
var nameRegexp *regexp.Regexp

func init() {
	numRegexp = regexp.MustCompile(`^(([0-9]+\.?[0-9]*)|([0-9]+/[0-9]+))$`)
	nameRegexp = regexp.MustCompile(`^[\w?+\-/*]`)
}

func Parse(source io.Reader) Parser {
	p := &parser{
		lexer: lexer.Lex(source),
	}

	return p
}

type parser struct {
	lexer lexer.Lexer
}

func (p *parser) Next() (Tree, error) {
	token, err := p.lexer.Next()
	if err != nil {
		return nil, err
	}
	return p.follow(token)
}

// TODO: Add quote "'" support
func (p *parser) follow(t lexer.Token) (Tree, error) {
	switch t.Type() {
	case lexer.IdentToken:
		return parseIdent(t)
	case lexer.OpenToken:
		return parseSExpr(p, t)
	default:
		return nil, fmt.Errorf("unexpected token %#v", t)
	}
}

// TODO: add string literal support
func parseIdent(t lexer.Token) (Tree, error) {
	text := t.Text()
	switch {
	case numRegexp.MatchString(text):
		return NewNumber(text, t.Start(), t.End()), nil
	case nameRegexp.MatchString(text):
		return NewIdentifier(text, t.Start(), t.End()), nil
	default:
		return nil, fmt.Errorf("invalid identifer %s", text)
	}
}

// TODO: add cons '.' support
func parseSExpr(p *parser, open lexer.Token) (Tree, error) {
	start := open.Start()

	var elems []Tree
	for {
		next, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		switch next.Type() {
		case lexer.CloseToken:
			return NewSExpression(start, next.End(), elems...), nil
		default:
			tree, err := p.follow(next)
			if err != nil {
				return nil, err
			}
			elems = append(elems, tree)
		}
	}
}

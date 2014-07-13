package lexer

import (
	"strings"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type LexTest struct {
	Source string
	Stream []Token
}

var lexTests = []LexTest{
	{"a", []Token{base{typ: IdentToken, text: "a"}}},
	{"null?", []Token{base{typ: IdentToken, text: "null?"}}},
	{"**", []Token{base{typ: IdentToken, text: "**"}}},
	{"123khb1", []Token{base{typ: IdentToken, text: "123khb1"}}},
	{"fmt.Printf", []Token{base{typ: IdentToken, text: "fmt.Printf"}}},
	{"(+ 1 3)", []Token{
		base{typ: OpenToken, text: "("},
		base{typ: IdentToken, text: "+"},
		base{typ: IdentToken, text: "1"},
		base{typ: IdentToken, text: "3"},
		base{typ: CloseToken, text: ")"},
	}},

	{"'(1 2)", []Token{
		base{typ: QuoteToken, text: "'"},
		base{typ: OpenToken, text: "("},
		base{typ: IdentToken, text: "1"},
		base{typ: IdentToken, text: "2"},
		base{typ: CloseToken, text: ")"},
	}},

	{"(1 . 2)", []Token{
		base{typ: OpenToken, text: "("},
		base{typ: IdentToken, text: "1"},
		base{typ: ConsToken, text: "."},
		base{typ: IdentToken, text: "2"},
		base{typ: CloseToken, text: ")"},
	}},

	{`(+ 
		'(1 2) 
		(3 . 4)
	  )`, []Token{
		base{typ: OpenToken, text: "("},
		base{typ: IdentToken, text: "+"},
		base{typ: QuoteToken, text: "'"},
		base{typ: OpenToken, text: "("},
		base{typ: IdentToken, text: "1"},
		base{typ: IdentToken, text: "2"},
		base{typ: CloseToken, text: ")"},
		base{typ: OpenToken, text: "("},
		base{typ: IdentToken, text: "3"},
		base{typ: ConsToken, text: "."},
		base{typ: IdentToken, text: "4"},
		base{typ: CloseToken, text: ")"},
		base{typ: CloseToken, text: ")"},
	}},
}

func readAll(l Lexer) ([]Token, error) {
	var resp []Token
	for {
		x, err := l.Next()
		if err != nil {
			if err == Done {
				return resp, nil
			}
			return nil, err
		}
		resp = append(resp, x)
	}
}

func cmpToken(a, b Token) bool {
	return a.Type() == b.Type() && a.Text() == b.Text()
}

func TestLexSimple(t *testing.T) {
	Convey("Given Some Source Input", t, func() {
		for _, lt := range lexTests {
			l := Lex(strings.NewReader(lt.Source))
			stream, err := readAll(l)
			So(err, ShouldBeNil)

			for i := range lt.Stream {
				t.Log(stream[i])
				So(cmpToken(stream[i], lt.Stream[i]), ShouldBeTrue)
			}
		}
	})
}

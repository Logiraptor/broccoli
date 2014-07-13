package parser

import (
	"strings"
	"testing"

	"github.com/Logiraptor/broccoli/lexer"
	. "github.com/smartystreets/goconvey/convey"
)

type ParseTest struct {
	Source string
	Stream []Tree
}

var p lexer.Pos

var parseTests = []ParseTest{
	{"a", []Tree{NewIdentifier("a", p, p)}},
	{"fmt.Printf", []Tree{NewIdentifier("fmt.Printf", p, p)}},
	{"123", []Tree{NewNumber("123", p, p)}},
	{"123.123", []Tree{NewNumber("123.123", p, p)}},
	{"56/432", []Tree{NewNumber("56/432", p, p)}},
	{"()", []Tree{NewSExpression(p, p)}},
	{"(+ 1 2)", []Tree{NewSExpression(p, p, NewIdentifier("+", p, p), NewNumber("1", p, p), NewNumber("2", p, p))}},
	{"(+ (* 3 5) (- 9 3))", []Tree{
		NewSExpression(p, p,
			NewIdentifier("+", p, p),
			NewSExpression(p, p,
				NewIdentifier("*", p, p),
				NewNumber("3", p, p),
				NewNumber("5", p, p),
			),
			NewSExpression(p, p,
				NewIdentifier("-", p, p),
				NewNumber("9", p, p),
				NewNumber("3", p, p),
			),
		),
	},
	},
}

func readAll(p Parser) ([]Tree, error) {
	var resp []Tree
	for {
		t, err := p.Next()
		if err != nil {
			if err == lexer.Done {
				return resp, nil
			}
			return nil, err
		}
		resp = append(resp, t)

	}
}

func cmpTree(a, b Tree, t *testing.T) {
	switch v := a.(type) {
	case SExpression:
		if o, ok := b.(SExpression); ok {
			t.Logf("%#v %#v", v, o)
			So(len(v.Elems), ShouldEqual, len(o.Elems))
			for i := range v.Elems {
				cmpTree(v.Elems[i], o.Elems[i], t)
			}
		}
	case Identifier:
		if o, ok := b.(Identifier); ok {
			So(v.Name, ShouldEqual, o.Name)
		}
	case Number:
		if o, ok := b.(Number); ok {
			So(v.Val, ShouldEqual, o.Val)
		}
	}
	So(a.Type(), ShouldEqual, b.Type())
}

func TestParser(t *testing.T) {
	Convey("Given an input string", t, func() {
		for _, pt := range parseTests {
			p := Parse(strings.NewReader(pt.Source))
			stream, err := readAll(p)
			So(err, ShouldBeNil)
			t.Logf("%#v %#v", stream, pt.Stream)
			So(len(stream), ShouldEqual, len(pt.Stream))

			for i := range pt.Stream {
				t.Log(stream[i])
				cmpTree(pt.Stream[i], stream[i], t)
			}
		}
	})
}

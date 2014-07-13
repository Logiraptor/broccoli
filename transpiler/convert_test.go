package transpiler

import (
	"bytes"
	"go/printer"
	"go/token"
	"strings"
	"testing"

	"github.com/Logiraptor/broccoli/parser"
	. "github.com/smartystreets/goconvey/convey"
)

type ConvertTest struct {
	Source   string
	GoSource string
}

var convertTests = []ConvertTest{
	{"a", "a"},
	{"1", "1"},
	{"(foo 1 2)", "foo(1, 2)"},
	{"(+ 1 2)", "1 + 2"},
	{"(* 1 2)", "1 * 2"},
	{"(* (+ 1 2) (- 9 5))", "(1 + 2) * (9 - 5)"},
	{"(fmt.Print (* (+ 1 2) (- 9 5)))", "fmt.Print((1 + 2) * (9 - 5))"},
	{"(var n 0)", "var n = 0"},
	{"(set n (+ n 1))", "n = n + 1"},
	{"(fn (n) ((+ n 1)))", "func(n int) int {\n\treturn n + 1\n}"},
	{"(if (== n 0) ((set n (n + 1))))", "if (n == 0) { n = n + 1 }"},
}

func TestConvert(t *testing.T) {
	Convey("Given some source", t, func() {
		for _, ct := range convertTests {
			p := parser.Parse(strings.NewReader(ct.Source))
			sexpr, _ := p.Next()
			node := Convert(sexpr)

			var buf = new(bytes.Buffer)
			err := printer.Fprint(buf, token.NewFileSet(), node)
			if err != nil {
				t.Fatal(err.Error())
			}
			actual := string(buf.Bytes())

			t.Logf("Go: %s", ct.GoSource)
			t.Logf("Actual: %s", actual)
			t.Logf("Broccoli: %s", ct.Source)

			So(actual, ShouldEqual, ct.GoSource)
		}
	})
}

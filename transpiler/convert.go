package transpiler

import (
	"go/ast"
	"go/token"

	"github.com/Logiraptor/broccoli/parser"
)

func Convert(t parser.Tree) ast.Node {
	switch v := t.(type) {
	case parser.Identifier:
		return ast.NewIdent(v.Name)
	case parser.Number:
		return &ast.BasicLit{
			Kind:  token.INT,
			Value: v.Val,
		}
	case parser.SExpression:
		fun := v.Elems[0]
		return callFunc(fun, v.Elems[1:]...)
	default:
		return nil
	}
}

func isOperator(name string) bool {
	return name == "+" ||
		name == "-" ||
		name == "*" ||
		name == "/"
}

func getOperator(name string) token.Token {
	switch name {
	case "+":
		return token.ADD
	case "-":
		return token.SUB
	case "*":
		return token.MUL
	case "/":
		return token.QUO
	default:
		return token.ILLEGAL
	}
}

func callFunc(fun parser.Tree, args ...parser.Tree) ast.Node {
	switch v := fun.(type) {
	case parser.Identifier:
		if isOperator(v.Name) {
			return &ast.BinaryExpr{
				X:  Convert(args[0]).(ast.Expr),
				Op: getOperator(v.Name),
				Y:  Convert(args[1]).(ast.Expr),
			}
		}
		switch v.Name {
		case "var":
			return &ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names: []*ast.Ident{
								ast.NewIdent(args[0].(parser.Identifier).Name),
							},
							Values: []ast.Expr{
								Convert(args[1]).(ast.Expr),
							},
						},
					},
				},
			}
		case "set":
			return &ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent(args[0].(parser.Identifier).Name)},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{Convert(args[1]).(ast.Expr)},
			}
		case "fn":
			var params []*ast.Ident
			argList := args[0].(parser.SExpression)
			for _, argName := range argList.Elems {
				params = append(params, ast.NewIdent(argName.(parser.Identifier).Name))
			}

			var body []ast.Stmt
			bodyList := args[1].(parser.SExpression)
			for _, line := range bodyList.Elems[:len(bodyList.Elems)-1] {
				body = append(body, Convert(line).(ast.Stmt))
			}
			body = append(body, &ast.ReturnStmt{
				Results: []ast.Expr{Convert(bodyList.Elems[len(bodyList.Elems)-1]).(ast.Expr)},
			})

			return &ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: params,
								Type:  ast.NewIdent("int"),
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Type: ast.NewIdent("int"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: body,
				},
			}
		default:
			return &ast.CallExpr{
				Fun:  ast.NewIdent(v.Name),
				Args: parseListExpr(args...),
			}
		}
	default:
		return nil
	}
}

func parseList(args ...parser.Tree) []ast.Node {
	var resp = make([]ast.Node, len(args))
	for i, a := range args {
		resp[i] = Convert(a)
	}
	return resp
}

func parseListExpr(args ...parser.Tree) []ast.Expr {
	var resp = make([]ast.Expr, len(args))
	for i, a := range args {
		resp[i] = Convert(a).(ast.Expr)
	}
	return resp
}

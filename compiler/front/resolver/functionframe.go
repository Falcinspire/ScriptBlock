package resolver

import "github.com/falcinspire/scriptblock/compiler/ast"

type FunctionFrame struct {
	Parameters []string
	Body       []ast.Statement
}

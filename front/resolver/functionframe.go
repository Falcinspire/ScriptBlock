package resolver

import "github.com/falcinspire/scriptblock/ast"

type FunctionFrame struct {
	Parameters []string
	Body       []ast.Statement
}

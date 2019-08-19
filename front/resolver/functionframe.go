package resolver

import "github.com/falcinspire/scriptblock/front/ast"

type FunctionFrame struct {
	Parameters []string
	Body       []ast.Statement
}

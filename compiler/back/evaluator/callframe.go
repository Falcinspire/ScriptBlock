package evaluator

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
)

type CallFrame struct {
	// this function
	Location *location.UnitLocation
	Name     string

	// which has internals
	Body       []ast.Statement
	Parameters []string
	Closes     []string

	// with these arguments
	Arguments []values.Value
	Captures  []values.Value
}

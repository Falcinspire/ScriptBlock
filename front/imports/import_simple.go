package imports

import (
	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/front/location"
)

func TakeImports(astree *ast.Unit) []*location.UnitLocation {
	imports := make([]*location.UnitLocation, len(astree.ImportLines))
	for i, importLine := range astree.ImportLines {
		imports[i] = location.NewUnitLocation(importLine.Module, importLine.Unit)
	}
	return imports
}

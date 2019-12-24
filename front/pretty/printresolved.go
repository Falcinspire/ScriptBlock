package pretty

import (
	"encoding/json"
	"fmt"

	"github.com/falcinspire/scriptblock/ast"
)

func PrintResolved(unit *ast.Unit) {
	imports := make([]*ResolvedImport, len(unit.ImportLines))
	for i, importLine := range unit.ImportLines {
		imports[i] = &ResolvedImport{Module: importLine.Module, Unit: importLine.Unit}
	}

	definitions := make([]ResolvedTopDefinition, len(unit.Definitions))
	for i, definition := range unit.Definitions {
		visitor := &ResolvedTopDefinitionVisitor{nil, nil}
		definition.Accept(visitor)
		definitions[i] = visitor.Result
	}

	jsonData, _ := json.Marshal(&ResolvedUnit{imports, definitions})
	fmt.Println(string(jsonData))
}

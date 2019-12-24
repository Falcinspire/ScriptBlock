package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/falcinspire/scriptblock/ast"
	"github.com/falcinspire/scriptblock/ast/symbol"
	"github.com/falcinspire/scriptblock/front/imports"
	"github.com/falcinspire/scriptblock/ast/location"
	"github.com/falcinspire/scriptblock/front/symbols"
)

type ResolveError struct {
	name     string
	metadata *ast.Metadata
}

func NewResolveError(symbol string, metadata *ast.Metadata) ResolveError {
	return ResolveError{symbol, metadata}
}
func (this ResolveError) Error() string {
	data, _ := json.Marshal(this.metadata)
	return fmt.Sprintf("Cannot resolve \"%s\", %s", this.name, string(data))
}

type ResolveEnvironment struct {
	linkedLocals *symbols.LinkedSymbolTable
	selfInternal symbols.SymbolTable
	selfExported symbols.SymbolTable
	imported     []symbols.SymbolTable
}

func CreateResolveEnvironment(unitLocation *location.UnitLocation, symbollibrary *symbols.SymbolLibrary, importbook imports.ImportBook) *ResolveEnvironment {
	internal := symbols.LookupSymbolTable(unitLocation.Module, unitLocation.Unit, symbollibrary.Internal)
	exported := symbols.LookupSymbolTable(unitLocation.Module, unitLocation.Unit, symbollibrary.Exported)
	usedImports := imports.LookupImportList(unitLocation.Module, unitLocation.Unit, importbook)
	importedTables := make([]symbols.SymbolTable, len(usedImports))
	for i, usedImport := range usedImports {
		importedTables[i] = symbols.LookupSymbolTable(usedImport.Module, usedImport.Unit, symbollibrary.Exported)
	}
	return NewResolveEnvironment(internal, exported, importedTables)
}

func NewResolveEnvironment(selfInternal, selfExported symbols.SymbolTable, imported []symbols.SymbolTable) *ResolveEnvironment {
	return &ResolveEnvironment{symbols.NewLinkedSymbolTable(), selfInternal, selfExported, imported}
}

func FindLocal(name string, env *ResolveEnvironment) (found *symbol.AddressBox, exists bool) {
	found, exists = env.linkedLocals.PeekTable()[name]
	return
}

func FindClosed(name string, env *ResolveEnvironment) (found *symbol.AddressBox, exists bool) {
	for i := env.linkedLocals.Length() - 2; i >= 0; i-- {
		closedFound, closedExists := env.linkedLocals.GetTableAt(i)[name]
		if closedExists {
			found = closedFound
			exists = true
			return
		}
	}
	return nil, false
}

func FindUnit(name string, env *ResolveEnvironment) (found *symbol.AddressBox, exists bool) {
	exported, exists := env.selfExported[name]
	if exists {
		return exported, true
	}

	internal, exists := env.selfInternal[name]
	if exists {
		return internal, true
	}

	return nil, false
}

func FindImported(name string, env *ResolveEnvironment) (found *symbol.AddressBox, exists bool) {
	for _, importedTable := range env.imported {
		imported, exists := importedTable[name]
		if exists {
			return imported, true
		}
	}
	return nil, false
}

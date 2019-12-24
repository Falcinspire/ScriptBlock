package valuepass

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
	"github.com/falcinspire/scriptblock/compiler/back/addressbook"
	"github.com/falcinspire/scriptblock/compiler/back/evaluator"
	"github.com/falcinspire/scriptblock/compiler/back/tags"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
)

// ValuePass performs all the necessary passes to evaluate and create .mcfunctions for the tree
func ValuePass(unit *ast.Unit, location *location.UnitLocation, valueLibrary *values.ValueLibrary, addressBook addressbook.AddressBook, tags map[string]tags.LocationList, modulePath string, output string) {
	data := &evaluator.EvaluateData{
		Location:     location,
		LocalValues:  nil,
		ValueLibrary: valueLibrary,
		AddressBook:  addressBook,
		CallStack:    evaluator.NewCallStack(),
		LoopInject:   evaluator.NewLoopInject(),
		ModulePath:   modulePath,
		Output:       output,
	}
	addressbook.InsertAddressTable(location.Module, location.Unit, addressbook.NewAddressTable(), addressBook)
	values.InsertValueTable(location.Module, location.Unit, values.NewValueTable(), valueLibrary.Exported)
	values.InsertValueTable(location.Module, location.Unit, values.NewValueTable(), valueLibrary.Internal)
	PassOne(unit, data)
	PassTwo(unit, data)
	PassThree(unit, data, tags)
}

package valuepass

import (
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/output"
	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/addressbook"
	"github.com/falcinspire/scriptblock/front/ast"
	"github.com/falcinspire/scriptblock/front/location"
)

// ValuePass performs all the necessary passes to evaluate and create .mcfunctions for the tree
func ValuePass(unit *ast.Unit, location *location.UnitLocation, valueLibrary *values.ValueLibrary, addressBook addressbook.AddressBook, tags map[string]tags.LocationList, output output.OutputDirectory) {
	data := &evaluator.EvaluateData{
		Location:     location,
		LocalValues:  nil,
		ValueLibrary: valueLibrary,
		AddressBook:  addressBook,
		LoopInject:   evaluator.NewLoopInject(),
		Output:       output,
	}
	addressbook.InsertAddressTable(location.Module, location.Unit, addressbook.NewAddressTable(), addressBook)
	values.InsertValueTable(location.Module, location.Unit, values.NewValueTable(), valueLibrary.Exported)
	values.InsertValueTable(location.Module, location.Unit, values.NewValueTable(), valueLibrary.Internal)
	PassOne(unit, data)
	PassTwo(unit, data)
	PassThree(unit, data, tags)
}

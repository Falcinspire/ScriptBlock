package addressbook

import (
	"github.com/falcinspire/scriptblock/compiler/ast"
)

// AddressBook stores the actual function code for each unit.
type AddressBook map[string]addressUnitBook

func NewAddressBook() AddressBook {
	return make(AddressBook)
}

type addressUnitBook map[string]*AddressTable

func newAddressUnitBook() addressUnitBook {
	return make(addressUnitBook)
}

func InsertAddressTable(module, unit string, table *AddressTable, book AddressBook) {
	_, exists := book[module]
	if !exists {
		book[module] = newAddressUnitBook()
	}

	book[module][unit] = table
}

func LookupAddressTable(module, unit string, book AddressBook) *AddressTable {
	return book[module][unit]
}

type AddressTable struct {
	Functors map[string]*ast.TemplateDefinition
}

func NewAddressTable() *AddressTable {
	return &AddressTable{make(map[string]*ast.TemplateDefinition)}
}

func AddressTemplate(module string, unit string, name string, book AddressBook) *ast.TemplateDefinition {
	return LookupAddressTable(module, unit, book).Functors[name]
}

func InsertFunctorAddress(closure *ast.TemplateDefinition, module string, unit string, name string, book AddressBook) {
	LookupAddressTable(module, unit, book).Functors[name] = closure
}

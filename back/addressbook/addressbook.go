package addressbook

import (
	"github.com/falcinspire/scriptblock/ast"
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
	Templates map[string]*ast.TemplateDefinition
	Closures  map[string]*ast.ClosureDefinition
}

func NewAddressTable() *AddressTable {
	return &AddressTable{make(map[string]*ast.TemplateDefinition), make(map[string]*ast.ClosureDefinition)}
}

func AddressTemplate(module string, unit string, name string, book AddressBook) *ast.TemplateDefinition {
	return LookupAddressTable(module, unit, book).Templates[name]
}

func AddressClosure(module string, unit string, name string, book AddressBook) *ast.ClosureDefinition {
	return LookupAddressTable(module, unit, book).Closures[name]
}

func InsertTemplateAddress(template *ast.TemplateDefinition, module string, unit string, name string, book AddressBook) {
	LookupAddressTable(module, unit, book).Templates[name] = template
}

func InsertClosureAddress(closure *ast.ClosureDefinition, module string, unit string, name string, book AddressBook) {
	LookupAddressTable(module, unit, book).Closures[name] = closure
}

package addressbook

import (
	"github.com/falcinspire/scriptblock/front/ast"
)

type AddressBook map[string]AddressUnitBook

func NewAddressBook() AddressBook {
	return make(AddressBook)
}

type AddressUnitBook map[string]*AddressTable

func NewAddressUnitBook() AddressUnitBook {
	return make(AddressUnitBook)
}

func InsertAddressTable(module, unit string, table *AddressTable, book AddressBook) {
	_, exists := book[module]
	if !exists {
		book[module] = NewAddressUnitBook()
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

package imports

import "github.com/falcinspire/scriptblock/front/location"

type ImportBook map[string]ImportUnitBook

func NewImportBook() ImportBook {
	return make(ImportBook)
}

func LookupImportList(module string, unit string, book ImportBook) ImportList {
	return book[module][unit]
}
func InsertImportList(module string, unit string, list ImportList, book ImportBook) {
	_, exists := book[module]
	if !exists {
		book[module] = make(ImportUnitBook)
	}

	book[module][unit] = list
}

type ImportUnitBook map[string]ImportList

type ImportList []*location.UnitLocation

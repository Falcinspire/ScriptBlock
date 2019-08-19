package values

type ValueLibrary struct {
	Exported ValueBook
	Internal ValueBook
}

func NewValueLibrary() *ValueLibrary {
	return &ValueLibrary{NewValueBook(), NewValueBook()}
}

func InsertValueTable(module, unit string, table ValueTable, book ValueBook) {
	_, exists := book[module]
	if !exists {
		book[module] = NewValueUnitBook()
	}

	book[module][unit] = table
}

func LookupValueTable(module, unit string, book ValueBook) ValueTable {
	return book[module][unit]
}

type ValueBook map[string]ValueUnitBook

func NewValueBook() ValueBook {
	return make(ValueBook)
}

type ValueUnitBook map[string]ValueTable

func NewValueUnitBook() ValueUnitBook {
	return make(ValueUnitBook)
}

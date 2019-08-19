package symbols

type SymbolLibrary struct {
	Exported SymbolBook
	Internal SymbolBook
}

func NewSymbolLibrary() SymbolLibrary {
	return SymbolLibrary{make(SymbolBook), make(SymbolBook)}
}

type SymbolBook map[string]SymbolUnitBook

func LookupSymbolTable(module string, unit string, symbolbook SymbolBook) SymbolTable {
	return symbolbook[module][unit]
}
func InsertSymbolTable(module string, unit string, symboltable SymbolTable, symbolbook SymbolBook) {
	_, exists := symbolbook[module]
	if !exists {
		symbolbook[module] = make(SymbolUnitBook)
	}

	symbolbook[module][unit] = symboltable
}

type SymbolUnitBook map[string]SymbolTable

func NewSymbolUnitBook() SymbolUnitBook {
	return make(SymbolUnitBook)
}

func NewSymbolBook() SymbolBook {
	return make(SymbolBook)
}

package symbols

type SymbolTable map[string]*AddressBox

func NewSymbolTable() SymbolTable {
	return make(SymbolTable)
}

type LocalSymbolTable map[string]*AddressBox

func NewLocalSymbolTable() LocalSymbolTable {
	return make(LocalSymbolTable)
}

type LinkedSymbolTable struct {
	linked []LocalSymbolTable
}

func NewLinkedSymbolTable() *LinkedSymbolTable {
	return &LinkedSymbolTable{[]LocalSymbolTable{}}
}

func (this *LinkedSymbolTable) PeekTable() LocalSymbolTable {
	return this.linked[len(this.linked)-1]
}

func (this *LinkedSymbolTable) GetTableAt(i int) LocalSymbolTable {
	return this.linked[i]
}

func (this *LinkedSymbolTable) PushTable(table LocalSymbolTable) {
	this.linked = append(this.linked, table)
}

func (this *LinkedSymbolTable) PopTable() {
	this.linked = this.linked[0 : len(this.linked)-1]
}

func (this *LinkedSymbolTable) Length() int {
	return len(this.linked)
}

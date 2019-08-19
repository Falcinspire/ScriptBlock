package resolver

type FreeTable []string

func NewFreeTable() FreeTable {
	return make(FreeTable, 0)
}

type FreeTableStack struct {
	stack []FreeTable
}

func NewFreeTableStack() *FreeTableStack {
	return &FreeTableStack{make([]FreeTable, 0)}
}

func (this *FreeTableStack) AppendToTop(identifier string) {
	this.stack[len(this.stack)-1] = append(this.stack[len(this.stack)-1], identifier)
}

func (this *FreeTableStack) PeekFreeTable() FreeTable {
	return this.stack[len(this.stack)-1]
}

func (this *FreeTableStack) PushFreeTable() {
	this.stack = append(this.stack, NewFreeTable())
}

func (this *FreeTableStack) PopFreeTable() FreeTable {
	remember := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	return remember
}

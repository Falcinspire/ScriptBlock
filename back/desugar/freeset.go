package desugar

import (
	"github.com/falcinspire/scriptblock/ast/symbol"
)

type freeVariableSet struct {
	list []*symbol.ParameterAddress
	set  map[string]bool
}

func newFreeVariableSet() *freeVariableSet {
	return &freeVariableSet{[]*symbol.ParameterAddress{}, map[string]bool{}}
}

func AddToFreeSet(address *symbol.ParameterAddress, set *freeVariableSet) {
	_, exists := set.set[address.Name]
	if !exists {
		set.set[address.Name] = true
		set.list = append(set.list, address)
	}
}

func ListFreeSet(set *freeVariableSet) []*symbol.ParameterAddress {
	return set.list
}

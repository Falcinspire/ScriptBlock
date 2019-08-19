package desugar

import (
	"github.com/falcinspire/scriptblock/front/symbols"
)

type freeVariableSet struct {
	list []*symbols.ParameterAddress
	set  map[string]bool
}

func newFreeVariableSet() *freeVariableSet {
	return &freeVariableSet{[]*symbols.ParameterAddress{}, map[string]bool{}}
}

func AddToFreeSet(address *symbols.ParameterAddress, set *freeVariableSet) {
	_, exists := set.set[address.Name]
	if !exists {
		set.set[address.Name] = true
		set.list = append(set.list, address)
	}
}

func ListFreeSet(set *freeVariableSet) []*symbols.ParameterAddress {
	return set.list
}

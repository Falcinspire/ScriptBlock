package evaluator

import (
	"github.com/falcinspire/scriptblock/back/addressbook"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/location"
)

type EvaluateData struct {
	Location     *location.UnitLocation
	LocalValues  *values.LocalValueTable
	ValueLibrary *values.ValueLibrary
	AddressBook  addressbook.AddressBook
	CallStack    *CallStack
	LoopInject   *LoopInject
	ModulePath   string
	Output       string
}

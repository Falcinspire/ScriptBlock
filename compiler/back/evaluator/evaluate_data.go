package evaluator

import (
	"github.com/falcinspire/scriptblock/compiler/back/addressbook"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
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

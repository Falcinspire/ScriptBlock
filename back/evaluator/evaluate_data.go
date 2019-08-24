package evaluator

import (
	"github.com/falcinspire/scriptblock/back/output"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/addressbook"
	"github.com/falcinspire/scriptblock/front/location"
)

type EvaluateData struct {
	Location     *location.UnitLocation
	LocalValues  *values.LocalValueTable
	ValueLibrary *values.ValueLibrary
	AddressBook  addressbook.AddressBook
	LoopInject   *LoopInject
	ModulePath   string
	Output       output.OutputDirectory
}

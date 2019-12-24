package maintool

import (
	"github.com/falcinspire/scriptblock/compiler/back/addressbook"
	"github.com/falcinspire/scriptblock/compiler/back/desugar"
	"github.com/falcinspire/scriptblock/compiler/back/tags"
	"github.com/falcinspire/scriptblock/compiler/back/valuepass"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/falcinspire/scriptblock/compiler/front/astbook"
	"github.com/falcinspire/scriptblock/compiler/front/imports"
	"github.com/falcinspire/scriptblock/compiler/ast/location"
	"github.com/falcinspire/scriptblock/compiler/front/resolver"
	"github.com/falcinspire/scriptblock/compiler/front/symbolgen"
	"github.com/falcinspire/scriptblock/compiler/front/symbols"
)

func DoUnitFront(unitLocation *location.UnitLocation, astbooko astbook.AstBook, importbook imports.ImportBook, symbollibrary *symbols.SymbolLibrary) {
	astree := astbook.LookupAst(unitLocation, astbooko)
	symbolgen.SymbolsPass(astree, unitLocation, symbollibrary)
	resolver.ResolvePass(astree, unitLocation, symbollibrary, importbook)
}

func DoUnitBack(unitLocation *location.UnitLocation, astbooko astbook.AstBook, valuelibrary *values.ValueLibrary, addressbook addressbook.AddressBook, theTags map[string]tags.LocationList, modulePath string, output string) {
	astree := astbook.LookupAst(unitLocation, astbooko)
	desugar.Unit(astree, unitLocation)
	valuepass.ValuePass(astree, unitLocation, valuelibrary, addressbook, theTags, modulePath, output)
}

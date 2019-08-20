package maintool

import (
	"github.com/falcinspire/scriptblock/back/desugar"
	"github.com/falcinspire/scriptblock/back/evaluator"
	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/back/valuepass"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/front/addressbook"
	"github.com/falcinspire/scriptblock/front/astbook"
	"github.com/falcinspire/scriptblock/front/imports"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/resolver"
	"github.com/falcinspire/scriptblock/front/symbolgen"
	"github.com/falcinspire/scriptblock/front/symbols"
)

func DoUnit(unitLocation *location.UnitLocation, astbooko astbook.AstBook, importbook imports.ImportBook, symbollibrary symbols.SymbolLibrary, valuelibrary *values.ValueLibrary, addressbook addressbook.AddressBook, theTags map[string]tags.LocationList, output evaluator.OutputDirectory) {
	// front end
	astree := astbook.LookupAst(unitLocation, astbooko)
	symbolgen.SymbolsPass(astree, unitLocation, symbollibrary)
	resolver.ResolvePass(astree, unitLocation, symbollibrary, importbook)

	// back end
	desugar.DesugarUnit(astree, unitLocation)
	valuepass.ValuePass(astree, unitLocation, valuelibrary, addressbook, theTags, output)
}

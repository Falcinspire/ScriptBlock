package maintool

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/front/astgen"
	"github.com/falcinspire/scriptblock/front/parser"

	"github.com/falcinspire/scriptblock/back/addressbook"
	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/dependency"
	"github.com/falcinspire/scriptblock/front/astbook"
	"github.com/falcinspire/scriptblock/front/imports"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
)

type UnitLocationPath struct {
	location *location.UnitLocation
	filepath string
}

func DoModule(moduleQualified string, scriptblockHome string, output string, astbooko astbook.AstBook, importbooko imports.ImportBook, symbollibrary *symbols.SymbolLibrary, valuelibrary *values.ValueLibrary, addressbooko addressbook.AddressBook) {

	// this could be dangerous if given wrong directory
	// if _, err := os.Stat(output); !os.IsNotExist(err) {
	// 	os.RemoveAll(output)
	// 	os.MkdirAll(output, os.ModePerm)
	// }

	modulePath := filepath.Join(scriptblockHome, moduleQualified)

	logrus.WithFields(logrus.Fields{
		"module": modulePath,
	}).Info("compiling module")

	files, err := ioutil.ReadDir(modulePath)
	if err != nil {
		panic(err)
	}

	moduleName := filepath.Base(filepath.Dir(modulePath)) // module/version
	locations := registerLocations(files, moduleName)
	parseAllToAsts(locations, moduleName, modulePath, astbooko)
	takeImportsFromAsts(locations, moduleName, astbooko, importbooko)
	order := makeDependencyOrder(locations, moduleName, importbooko)

	logrus.WithFields(logrus.Fields{
		"order": order,
	}).Info("dependency order produced")

	RunFrontEnd(order, moduleName, astbooko, importbooko, symbollibrary)

	theTags := make(map[string]tags.LocationList)

	RunBackEnd(order, moduleName, astbooko, valuelibrary, addressbooko, theTags, modulePath, output)

	tags.WriteAllTagsToFiles(theTags, output)
}

func registerLocations(files []os.FileInfo, moduleName string) []string {
	locations := make([]string, 0)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sb" {
			unitName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			locations = append(locations, unitName)

			logrus.WithFields(logrus.Fields{
				"module": moduleName,
				"unit":   unitName,
			}).Info("identified unit")
		} else {
			logrus.WithFields(logrus.Fields{
				"module": moduleName,
				"file":   file.Name(),
			}).Info("skipping non-unit")
		}
	}

	return locations
}

func parseAllToAsts(units []string, moduleName string, modulePath string, astbooko astbook.AstBook) {
	for _, unitName := range units {
		filelocation := filepath.Join(modulePath, fmt.Sprintf("%s.sb", unitName))
		pstree := parser.Parse(filelocation)
		astree := astgen.PSTtoAST(pstree)
		astbook.InsertAst(&location.UnitLocation{Module: moduleName, Unit: unitName}, astree, astbooko)

		logrus.WithFields(logrus.Fields{
			"module": moduleName,
			"unit":   unitName,
			"path":   filelocation,
		}).Info("ast produced")
	}
}

func takeImportsFromAsts(units []string, module string, astbooko astbook.AstBook, importbooko imports.ImportBook) {
	for _, unitName := range units {
		astree := astbook.LookupAst(&location.UnitLocation{Module: module, Unit: unitName}, astbooko)
		importList := imports.TakeImports(astree)
		imports.InsertImportList(module, unitName, importList, importbooko)

		importListString := make([]string, len(importList))
		for i, importLine := range importList {
			importListString[i] = fmt.Sprintf("%s:%s", importLine.Module, importLine.Unit)
		}
		logrus.WithFields(logrus.Fields{
			"module":  module,
			"unit":    unitName,
			"imports": importListString,
		}).Info("imports taken")
	}
}

func makeDependencyOrder(units []string, module string, importbooko imports.ImportBook) []string {

	logrus.WithFields(logrus.Fields{
		"input": units,
	})

	toID := make(map[string]int)
	for unitID, unitName := range units {
		toID[unitName] = unitID
	}

	dependencyGraph := dependency.NewDependencyGraph()
	for i := 0; i < len(units); i++ {
		dependency.AddVertex(dependencyGraph)
	}

	// connect nodes
	for unitID, unitName := range units {
		importList := imports.LookupImportList(module, unitName, importbooko)
		for _, importLine := range importList {
			if importLine.Module == module {
				dependency.AddDependency(unitID, toID[importLine.Unit], dependencyGraph)
			}
		}
	}

	idOrder, circular := dependency.MakeDependencyOrder(dependencyGraph)
	if circular {
		panic(fmt.Errorf("circular reference"))
	}

	order := make([]string, len(units))
	for index, id := range idOrder {
		order[index] = units[id]
	}

	return order
}

func RunFrontEnd(order []string, module string, astbooko astbook.AstBook, importbooko imports.ImportBook, symbolLibrary *symbols.SymbolLibrary) {
	for _, unitName := range order {
		DoUnitFront(&location.UnitLocation{Module: module, Unit: unitName}, astbooko, importbooko, symbolLibrary)

		logrus.WithFields(logrus.Fields{
			"module": module,
			"unit":   unitName,
		}).Info("front end run on unit")
	}
}

func RunBackEnd(order []string, module string, astbooko astbook.AstBook, valueLibrary *values.ValueLibrary, addressbooko addressbook.AddressBook, theTags map[string]tags.LocationList, modulePath string, output string) {
	for _, unitName := range order {
		DoUnitBack(&location.UnitLocation{Module: module, Unit: unitName}, astbooko, valueLibrary, addressbooko, theTags, modulePath, output)

		logrus.WithFields(logrus.Fields{
			"module": module,
			"unit":   unitName,
		}).Info("back end run on unit")
	}
}

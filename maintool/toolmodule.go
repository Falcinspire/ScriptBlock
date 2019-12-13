package maintool

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/back/output"
	"github.com/falcinspire/scriptblock/back/tags"
	"github.com/falcinspire/scriptblock/front/astgen"
	"github.com/falcinspire/scriptblock/front/parser"

	"github.com/falcinspire/scriptblock/back/values"
	"github.com/falcinspire/scriptblock/dependency"
	"github.com/falcinspire/scriptblock/front/addressbook"
	"github.com/falcinspire/scriptblock/front/astbook"
	"github.com/falcinspire/scriptblock/front/imports"
	"github.com/falcinspire/scriptblock/front/location"
	"github.com/falcinspire/scriptblock/front/symbols"
)

type UnitLocationPath struct {
	location *location.UnitLocation
	filepath string
}

func DoModule(modulePath string, moduleQualified, output output.OutputDirectory) {

	// this could be dangerous if given wrong directory
	// if _, err := os.Stat(output); !os.IsNotExist(err) {
	// 	os.RemoveAll(output)
	// 	os.MkdirAll(output, os.ModePerm)
	// }

	logrus.WithFields(logrus.Fields{
		"module": modulePath,
	}).Info("compiling module")

	files, err := ioutil.ReadDir(modulePath)
	if err != nil {
		panic(err)
	}

	moduleName := filepath.Base(filepath.Dir(modulePath)) // module/version
	locations := registerLocations(files, moduleName)
	astbooko := parseAllToAsts(locations, moduleName, modulePath)
	importbooko := takeImportsFromAsts(locations, moduleName, astbooko)
	symbollibrary := symbols.NewSymbolLibrary()
	order := makeDependencyOrder(locations, moduleName, importbooko)

	logrus.WithFields(logrus.Fields{
		"order": order,
	}).Info("dependency order produced")

	RunFrontEnd(order, moduleName, astbooko, importbooko, symbollibrary)

	valuelibrary := values.NewValueLibrary()
	addressbooko := addressbook.NewAddressBook()

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

func parseAllToAsts(units []string, moduleName string, modulePath string) astbook.AstBook {
	astbooko := astbook.NewAstBook()
	for _, unitName := range units {
		filelocation := filepath.Join(modulePath, fmt.Sprintf("%s.sb", unitName))
		println(filelocation)
		pstree := parser.Parse(filelocation)
		astree := astgen.PSTtoAST(pstree)
		astbook.InsertAst(&location.UnitLocation{moduleName, unitName}, astree, astbooko)

		logrus.WithFields(logrus.Fields{
			"module": moduleName,
			"unit":   unitName,
			"path":   filelocation,
		}).Info("ast produced")
	}
	return astbooko
}

func takeImportsFromAsts(units []string, module string, astbooko astbook.AstBook) imports.ImportBook {
	importbooko := imports.NewImportBook()
	for _, unitName := range units {
		astree := astbook.LookupAst(&location.UnitLocation{module, unitName}, astbooko)
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
	return importbooko
}

func makeDependencyOrder(units []string, module string, importbooko imports.ImportBook) []string {

	logrus.WithFields(logrus.Fields{
		"input": units,
	})

	toId := make(map[string]int)
	for unitId, unitName := range units {
		toId[unitName] = unitId
	}

	dependencyGraph := dependency.NewDependencyGraph(len(units))

	// connect nodes
	for unitId, unitName := range units {
		importList := imports.LookupImportList(module, unitName, importbooko)
		for _, importLine := range importList {
			if importLine.Module == module {
				dependency.AddDependency(unitId, toId[importLine.Unit], dependencyGraph)
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
		DoUnitFront(&location.UnitLocation{module, unitName}, astbooko, importbooko, symbolLibrary)

		logrus.WithFields(logrus.Fields{
			"module": module,
			"unit":   unitName,
		}).Info("front end run on unit")
	}
}

func RunBackEnd(order []string, module string, astbooko astbook.AstBook, valueLibrary *values.ValueLibrary, addressbooko addressbook.AddressBook, theTags map[string]tags.LocationList, modulePath string, output output.OutputDirectory) {
	for _, unitName := range order {
		DoUnitBack(&location.UnitLocation{module, unitName}, astbooko, valueLibrary, addressbooko, theTags, modulePath, output)

		logrus.WithFields(logrus.Fields{
			"module": module,
			"unit":   unitName,
		}).Info("back end run on unit")
	}
}

package maintool

import (
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

func DoModule(path string, output output.OutputDirectory) {

	// this could be dangerous if given wrong directory
	// if _, err := os.Stat(output); !os.IsNotExist(err) {
	// 	os.RemoveAll(output)
	// 	os.MkdirAll(output, os.ModePerm)
	// }

	moduleFolder := filepath.Base(path)

	logrus.WithFields(logrus.Fields{
		"module": moduleFolder,
	}).Info("compiling module")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	locations := registerLocations(files, moduleFolder)
	astbooko := parseAllToAsts(locations)
	importbooko := takeImportsFromAsts(astbooko, locations)
	symbollibrary := symbols.NewSymbolLibrary()
	order := makeDependencyOrder(locations, importbooko)

	logrus.WithFields(logrus.Fields{
		"order": order,
	}).Info("dependency order produced")

	RunFrontEnd(order, astbooko, importbooko, symbollibrary)

	valuelibrary := values.NewValueLibrary()
	addressbooko := addressbook.NewAddressBook()

	theTags := make(map[string]tags.LocationList)

	RunBackEnd(order, astbooko, valuelibrary, addressbooko, theTags, path, output)

	tags.WriteAllTagsToFiles(theTags, output)
}

func registerLocations(files []os.FileInfo, moduleFolder string) []*location.UnitLocation {
	locations := make([]*location.UnitLocation, 0)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sb" {
			unitName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			locations = append(locations, location.NewUnitLocation(moduleFolder, unitName))

			logrus.WithFields(logrus.Fields{
				"module": moduleFolder,
				"unit":   unitName,
			}).Info("identified unit")
		} else {
			logrus.WithFields(logrus.Fields{
				"module": moduleFolder,
				"unit":   file.Name(),
			}).Info("skipping non-unit")
		}
	}

	return locations
}

func parseAllToAsts(locations []*location.UnitLocation) astbook.AstBook {
	astbooko := astbook.NewAstBook()
	for _, unitlocation := range locations {
		pstree := parser.Parse(unitlocation)
		astree := astgen.PSTtoAST(pstree)
		astbook.InsertAst(unitlocation, astree, astbooko)

		logrus.WithFields(logrus.Fields{
			"module": unitlocation.Module,
			"unit":   unitlocation.Unit,
		}).Info("ast produced")
	}
	return astbooko
}

func takeImportsFromAsts(astbooko astbook.AstBook, locations []*location.UnitLocation) imports.ImportBook {
	importbooko := imports.NewImportBook()
	for _, unitlocation := range locations {
		astree := astbook.LookupAst(unitlocation, astbooko)
		importList := imports.TakeImports(astree)
		imports.InsertImportList(unitlocation.Module, unitlocation.Unit, importList, importbooko)

		importListString := make([]string, len(importList))
		for i, importLine := range importList {
			importListString[i] = location.InformalPath(importLine)
		}
		logrus.WithFields(logrus.Fields{
			"module":  unitlocation.Module,
			"unit":    unitlocation.Unit,
			"imports": importListString,
		}).Info("imports taken")
	}
	return importbooko
}

func makeDependencyOrder(locations []*location.UnitLocation, importbooko imports.ImportBook) []string {
	dependencyGraph := dependency.NewDependencyGraph()
	// insert nodes
	for _, unitlocation := range locations {
		informalPath := location.InformalPath(unitlocation)
		dependency.InsertNode(informalPath, dependencyGraph)
	}
	// connect nodes
	for _, unitLocation := range locations {
		importList := imports.LookupImportList(unitLocation.Module, unitLocation.Unit, importbooko)
		informalDependent := location.InformalPath(unitLocation)
		for _, importLine := range importList {
			informalDependency := location.InformalPath(importLine)
			dependency.AddDependency(informalDependent, informalDependency, dependencyGraph)
		}
	}

	return dependency.MakeDependencyOrder(location.InformalPath(locations[0]), dependencyGraph)
}

func RunFrontEnd(order []string, astbooko astbook.AstBook, importbooko imports.ImportBook, symbolLibrary *symbols.SymbolLibrary) {
	for _, informalLocation := range order {
		unitLocation := location.LocationFromInformal(informalLocation)
		DoUnitFront(unitLocation, astbooko, importbooko, symbolLibrary)

		logrus.WithFields(logrus.Fields{
			"module": unitLocation.Module,
			"unit":   unitLocation.Unit,
		}).Info("front end run on unit")
	}
}

func RunBackEnd(order []string, astbooko astbook.AstBook, valueLibrary *values.ValueLibrary, addressbooko addressbook.AddressBook, theTags map[string]tags.LocationList, modulePath string, output output.OutputDirectory) {
	for _, informalLocation := range order {
		unitLocation := location.LocationFromInformal(informalLocation)
		DoUnitBack(unitLocation, astbooko, valueLibrary, addressbooko, theTags, modulePath, output)

		logrus.WithFields(logrus.Fields{
			"module": unitLocation.Module,
			"unit":   unitLocation.Unit,
		}).Info("back end run on unit")
	}
}

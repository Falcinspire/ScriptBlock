package maintool

import (
	"fmt"
	"path/filepath"

	"github.com/falcinspire/scriptblock/compiler/front/astbook"

	"github.com/falcinspire/scriptblock/compiler/back/addressbook"
	"github.com/falcinspire/scriptblock/compiler/back/values"
	"github.com/falcinspire/scriptblock/compiler/front/imports"
	"github.com/falcinspire/scriptblock/compiler/front/symbols"
	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/compiler/dependency"
	"github.com/falcinspire/scriptblock/environment"
	"github.com/falcinspire/scriptblock/home"
)

func DoProject(sourceModuleQualified string, output string) {
	scriptblockHome := environment.GetHomeSourcePath()

	order := makeModuleDependencyOrder(sourceModuleQualified, scriptblockHome)

	astbooko := astbook.NewAstBook()
	importbooko := imports.NewImportBook()
	symbollibrary := symbols.NewSymbolLibrary()
	valuelibrary := values.NewValueLibrary()
	addressbooko := addressbook.NewAddressBook()

	logrus.WithFields(logrus.Fields{
		"order": order,
	}).Info("module dependency order")

	for _, moduleQualified := range order {
		DoModule(moduleQualified, scriptblockHome, output, astbooko, importbooko, symbollibrary, valuelibrary, addressbooko)
	}
}

//TODO test multiple modules
func makeModuleDependencyOrder(moduleQualified string, scriptblockHome string) []string {
	graph := dependency.NewDependencyGraph()

	moduleToID := make(map[string]int)
	idToModule := make(map[int]string)

	addModuleDependencies(moduleQualified, scriptblockHome, graph, moduleToID, idToModule)
	orderID, circular := dependency.MakeDependencyOrder(graph)
	if circular {
		panic(fmt.Errorf("circular module dependency"))
	}

	order := make([]string, len(graph.Nodes))
	for index, i := range orderID {
		order[index] = idToModule[i]
	}

	return order
}

func addModuleDependencies(moduleQualified string, scriptblockHome string, graph *dependency.DependencyGraph, moduleToID map[string]int, idToModule map[int]string) int {
	modulePath := filepath.Join(scriptblockHome, moduleQualified)

	id, exists := moduleToID[moduleQualified]
	if exists {
		return id
	}

	id = dependency.AddVertex(graph)
	idToModule[id] = moduleQualified

	config := home.ReadModuleFile(filepath.Join(modulePath, "module.yaml"))
	for _, depends := range config.Dependencies {
		dependsID := addModuleDependencies(filepath.Join(depends.Location, depends.Version), scriptblockHome, graph, moduleToID, idToModule)
		dependency.AddDependency(id, dependsID, graph)
	}

	return id
}

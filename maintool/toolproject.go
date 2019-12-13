package maintool

import (
	"fmt"
	"path/filepath"

	"github.com/falcinspire/scriptblock/back/output"
	"github.com/sirupsen/logrus"

	"github.com/falcinspire/scriptblock/dependency"
	"github.com/falcinspire/scriptblock/environment"
	"github.com/falcinspire/scriptblock/home"
)

func DoProject(sourceModuleQualified string, output output.OutputDirectory) {
	scriptblockHome := environment.GetHomeSourcePath()

	order := makeModuleDependencyOrder(sourceModuleQualified, scriptblockHome)

	logrus.WithFields(logrus.Fields{
		"order": order,
	}).Info("module dependency order")

	for _, moduleQualified := range order {
		DoModule(moduleQualified, scriptblockHome, output)
	}
}

//TODO test multiple modules
func makeModuleDependencyOrder(moduleQualified string, scriptblockHome string) []string {
	graph := dependency.NewDependencyGraph()

	moduleToId := make(map[string]int)
	idToModule := make(map[int]string)

	addModuleDependencies(moduleQualified, scriptblockHome, graph, moduleToId, idToModule)
	orderId, circular := dependency.MakeDependencyOrder(graph)
	if circular {
		panic(fmt.Errorf("circular module dependence"))
	}

	order := make([]string, len(graph.Nodes))
	for index, i := range orderId {
		order[index] = idToModule[i]
	}

	return order
}

func addModuleDependencies(moduleQualified string, scriptblockHome string, graph *dependency.DependencyGraph, moduleToId map[string]int, idToModule map[int]string) int {
	modulePath := filepath.Join(scriptblockHome, moduleQualified)

	id, exists := moduleToId[moduleQualified]
	if exists {
		return id
	}

	id = dependency.AddVertex(graph)
	idToModule[id] = moduleQualified

	config := home.ReadModuleFile(filepath.Join(modulePath, "module.yaml"))
	for _, depends := range config.Dependencies {
		dependsID := addModuleDependencies(depends.Location, scriptblockHome, graph, moduleToId, idToModule)
		dependency.AddDependency(id, dependsID, graph)
	}

	return id
}

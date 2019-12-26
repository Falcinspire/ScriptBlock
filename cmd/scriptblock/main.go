package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/falcinspire/scriptblock/compiler/maintool"
	"github.com/falcinspire/scriptblock/home"
)

func main() {
	modulePrefix := os.Args[1]
	version := os.Args[2]
	module := filepath.Join(modulePrefix, version)
	_, exists := home.FindModuleInHome(module) // TODO remove redundancy
	if !exists {
		panic(fmt.Errorf("no module by name %s", module))
	}
	outputPath := home.MakeModuleOutput(module)
	maintool.DoProject(module, outputPath)
}

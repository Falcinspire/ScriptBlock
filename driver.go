package main

import (
	"fmt"
	"os"

	"github.com/falcinspire/scriptblock/home"
	"github.com/falcinspire/scriptblock/maintool"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {

	// predator, _ := home.FindModuleInHome("predator")
	// data, _ := json.Marshal(moduledat.ReadModuleData(predator))
	// fmt.Println(string(data))

	module := os.Args[1]
	modulePath, exists := home.FindModuleInHome(module)
	if !exists {
		panic(fmt.Errorf("no module by name %s", module))
	}
	outputPath := home.MakeModuleOutput(module)
	maintool.DoModule(modulePath, module, outputPath)
}

package home

import (
	"os"
	"path/filepath"

	"github.com/falcinspire/scriptblock/environment"
)

func FindModuleInHome(name string) (path string, exists bool) {
	homePath := os.Getenv("ScriptBlockPath")
	modulePath := filepath.Join(homePath, "src", name)
	if _, err := os.Stat(modulePath); !os.IsNotExist(err) {
		return modulePath, true
	} else {
		return modulePath, false
	}
}

func MakeModuleOutput(name string) string {
	homePath := os.Getenv("ScriptBlockPath")
	modulePath := filepath.Join(homePath, "bin", name)
	return modulePath
}

func GetModuleConfig(data environment.ModuleDescription) *ModuleData {
	configPath := filepath.Join(environment.GetModulePath(data), "module.yaml")
	return ReadModuleFile(configPath)
}

func MakeModulePath(data environment.ModuleDescription) {
	modulePath := filepath.Join(environment.GetHomeSourcePath(), data.Location, data.Version)
	os.MkdirAll(modulePath, os.ModePerm)
}

// does not include version
func MakeAbreviatedModulePath(data environment.ModuleDescription) {
	modulePath := filepath.Join(environment.GetHomeSourcePath(), data.Location)
	os.MkdirAll(modulePath, os.ModePerm)
}

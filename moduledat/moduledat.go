package moduledat

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ModuleData struct {
	Name         string   `yaml:"name"`
	Description  string   `yaml:"description"`
	Author       string   `yaml:"author"`
	Version      string   `yaml:"version"`
	Dependencies []string `yaml:"dependencies"`
}

func ReadModuleData(modulepath string) *ModuleData {
	filePath := filepath.Join(modulepath, "module.yaml")
	return ReadModuleFile(filePath)
}

func ReadModuleFile(path string) *ModuleData {
	data := ModuleData{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, &data)
	return &data
}

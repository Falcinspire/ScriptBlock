package home

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ModuleData struct {
	Name         string       `yaml:"name"`
	Description  string       `yaml:"description"`
	Author       string       `yaml:"author"`
	Dependencies []Dependency `yaml:"dependencies"`
}

type Dependency struct {
	Location string `yaml:"location"`
	Version  string `yaml:"version"`
	Name     string `yaml:"name"`
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

func WriteModuleData(data *ModuleData, path string) {
	bytes, _ := yaml.Marshal(data)
	ioutil.WriteFile(path, bytes, 0644)
}

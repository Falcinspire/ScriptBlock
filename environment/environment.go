package environment

import (
	"os"
	"path/filepath"
)

type ModuleDescription struct {
	Location string
	Version  string
}

func GetHomePath() string {
	return os.Getenv("ScriptBlockPath")
}

func GetHomeSourcePath() string {
	return filepath.Join(GetHomePath(), "src")
}

func GetHomeOutputPath() string {
	return filepath.Join(GetHomePath(), "bin")
}

func GetModulePath(data ModuleDescription) string {
	return filepath.Join(GetHomeSourcePath(), data.Location, data.Version)
}

func GetAbreviatedModulePath(data ModuleDescription) string {
	return filepath.Join(GetHomeSourcePath(), data.Location)
}

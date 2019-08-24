package home

import (
	"os"
	"path/filepath"
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

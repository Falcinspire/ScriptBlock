package dumper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func DumpFunction(module string, unit string, name string, lines []string, output string) {
	path := fmt.Sprintf("%s/%s/functions/%s/%s.mcfunction", output, module, unit, name)

	logrus.WithFields(logrus.Fields{
		"path":  path,
		"unit":  unit,
		"lines": len(lines),
	}).Info("dumping function")

	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, error := os.Create(path)
	if error != nil {
		panic(error)
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(line)
		file.WriteString("\n")
	}
}

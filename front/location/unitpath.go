package location

import (
	"fmt"
	"strings"
)

func InformalPath(location *UnitLocation) string {
	return fmt.Sprintf("%s:%s:%s", location.Location, location.Module, location.Unit)
}

func LocationFromInformal(path string) *UnitLocation {
	data := strings.Split(path, ":")
	return NewUnitLocation(data[0], data[1], data[2])
}

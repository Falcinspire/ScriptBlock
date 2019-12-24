package location

import (
	"fmt"
	"strings"
)

func InformalTagPath(location *TagLocation) string {
	return fmt.Sprintf("%s:%s", location.Namespace, location.Identity)
}

func TagLocationFromInformal(path string) *TagLocation {
	data := strings.Split(path, ":")
	return NewTagLocation(data[0], data[1])
}

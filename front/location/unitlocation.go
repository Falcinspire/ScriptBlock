package location

type UnitLocation struct {
	Location string
	Module   string
	Unit     string
}

func NewUnitLocation(location string, module string, unit string) *UnitLocation {
	return &UnitLocation{location, module, unit}
}

func Equals(a, b *UnitLocation) bool {
	return a.Location == b.Location && a.Module == b.Module && a.Unit == b.Unit
}

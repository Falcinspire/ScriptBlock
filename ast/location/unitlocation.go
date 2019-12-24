package location

type UnitLocation struct {
	Module string
	Unit   string
}

func NewUnitLocation(module string, unit string) *UnitLocation {
	return &UnitLocation{module, unit}
}

func Equals(a, b *UnitLocation) bool {
	return a.Module == b.Module && a.Unit == b.Unit
}

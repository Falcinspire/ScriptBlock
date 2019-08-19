package location

type TagLocation struct {
	Namespace string
	Identity  string
}

func NewTagLocation(namespace string, identity string) *TagLocation {
	return &TagLocation{namespace, identity}
}

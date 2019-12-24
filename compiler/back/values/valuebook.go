package values

type ValueTable map[string]Value

func NewValueTable() ValueTable {
	return make(ValueTable)
}

func LookupUnitValue(module string, unit string, function string, valuebook ValueTable) (value Value, exists bool) {
	value, exists = valuebook[function]
	return
}

type LocalValueTable struct {
	Parameters map[string]Value
	Captures   map[string]Value
}

func NewLocalValueTable(parameters map[string]Value, captures map[string]Value) *LocalValueTable {
	return &LocalValueTable{parameters, captures}
}

func LookupParameterValue(name string, local *LocalValueTable) (value Value, exists bool) {
	value, exists = local.Parameters[name]
	return
}

func LookupCaptureValue(name string, local *LocalValueTable) (value Value, exists bool) {
	value, exists = local.Captures[name]
	return
}

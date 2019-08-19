package nativefuncs

import (
	"fmt"
	"strings"

	"github.com/falcinspire/scriptblock/back/values"
)

// NativeFunction is a function that is run with arguments and returns a string
type NativeFunction func(args []values.Value) string

func numberAsInt(args []values.Value) string {
	valueAsInt := int(args[0].(*values.NumberValue).Value)
	return fmt.Sprintf("%d", valueAsInt)
}

func fromFile(args []values.Value) string {
	valueAsString := args[0].(*values.StringValue).Value
	return readResource(valueAsString)
}

func joinToSelector(args []values.Value) string {
	valuesAsStringArray := make([]string, len(args))
	for i, arrayValue := range args {
		valuesAsStringArray[i] = arrayValue.(*values.StringValue).Value
	}
	body := strings.Join(valuesAsStringArray, ",")
	return fmt.Sprintf("@e[%s]", body)
}

func joinStrings(args []values.Value) string {
	valuesAsStringArray := make([]string, len(args))
	for i, arrayValue := range args {
		valuesAsStringArray[i] = fmt.Sprintf("%s", arrayValue)
	}
	return strings.Join(valuesAsStringArray, "")
}

func giveQuotes(args []values.Value) string {
	valueAsString := args[0].(*values.StringValue).Value
	return fmt.Sprintf("\"%s\"", valueAsString)
}

// CreateNativeMap creates a map of names to native functions
func CreateNativeMap() map[string]NativeFunction {
	return map[string]NativeFunction{
		"int":      numberAsInt,
		"resource": fromFile,
		"selector": joinToSelector,
		"join":     joinStrings,
		"string":   giveQuotes,
	}
}

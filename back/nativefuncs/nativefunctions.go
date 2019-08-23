package nativefuncs

import (
	"fmt"
	"strings"

	"github.com/falcinspire/scriptblock/back/values"
)

// NativeFunction is a function that is run with arguments and returns a string
type NativeFunction func(args []values.Value) string

// numberAsInt converts the number given as the first argument into an integer
func numberAsInt(args []values.Value) string {
	valueAsInt := int(args[0].(*values.NumberValue).Value)
	return fmt.Sprintf("%d", valueAsInt)
}

// fromFile reads the file at the string path provided and returns its contents
// without whitespaces
func fromFile(args []values.Value) string {
	valueAsString := args[0].(*values.StringValue).Value
	return readResource(valueAsString)
}

// joinToSelector joins all the arguments as strings into a selector
func joinToSelector(args []values.Value) string {
	valuesAsStringArray := make([]string, len(args))
	for i, arrayValue := range args {
		valuesAsStringArray[i] = arrayValue.(*values.StringValue).Value
	}
	body := strings.Join(valuesAsStringArray, ",")
	return fmt.Sprintf("@e[%s]", body)
}

// joinStrings joins all the arguments as strings into a string.
// This is useful when things need to be joined without a space inbetween.
func joinStrings(args []values.Value) string {
	valuesAsStringArray := make([]string, len(args))
	for i, arrayValue := range args {
		valuesAsStringArray[i] = fmt.Sprintf("%s", arrayValue)
	}
	return strings.Join(valuesAsStringArray, "")
}

// giveQuotes puts quotes around the first argument provided.
// This makes the argument into a string expression.
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

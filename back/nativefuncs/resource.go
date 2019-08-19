package nativefuncs

import (
	"io/ioutil"
	"strings"
	"unicode"
)

// readResource reads the file at filepath into a no whitespace string
func readResource(filepath string) string {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)

	fileNoWs := spaceStringsBuilder(fileString)

	return fileNoWs
}

// https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
// author Tim Cooper
func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

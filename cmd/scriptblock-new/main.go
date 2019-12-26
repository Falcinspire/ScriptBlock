package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/falcinspire/scriptblock/environment"
	"github.com/falcinspire/scriptblock/home"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	name := prompt("Enter the name of your module: ", reader)
	description := prompt("Enter a one-line description: ", reader)
	version := prompt("Enter a version code: ", reader)
	author := prompt("Enter author name: ", reader)
	location := prompt("Enter the github url: ", reader)
	// ignore dependencies for now

	modulePath := filepath.Join(environment.GetHomeSourcePath(), location, version)
	fmt.Println(modulePath)
	os.MkdirAll(modulePath, os.ModePerm)

	dataPath := filepath.Join(modulePath, "module.yaml")
	fmt.Println(dataPath)
	home.WriteModuleData(&home.ModuleData{
		Name:         name,
		Description:  description,
		Author:       author,
		Dependencies: []home.Dependency{},
	}, dataPath)
}

// ref https://tutorialedge.net/golang/reading-console-input-golang/
func prompt(question string, reader *bufio.Reader) string {
	fmt.Print(question)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\r\n", "", -1)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/falcinspire/scriptblock/downloader"
	"github.com/falcinspire/scriptblock/environment"
	"github.com/falcinspire/scriptblock/home"
	"github.com/falcinspire/scriptblock/compiler/maintool"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	subcommand := os.Args[1]

	if subcommand == "build" {
		modulePrefix := os.Args[2]
		version := os.Args[3]
		module := filepath.Join(modulePrefix, version)
		_, exists := home.FindModuleInHome(module) // TODO remove redundancy
		if !exists {
			panic(fmt.Errorf("no module by name %s", module))
		}
		outputPath := home.MakeModuleOutput(module)
		maintool.DoProject(module, outputPath)
	} else if subcommand == "home" {
		fmt.Println(environment.GetHomePath())
	} else if subcommand == "config" {
		location := os.Args[2]
		version := os.Args[3]
		data := home.GetModuleConfig(environment.ModuleDescription{
			Location: location,
			Version:  version,
		})
		fmt.Println(data.Name)
		fmt.Println(data.Description)
		fmt.Println(data.Author)
		fmt.Println(data.Dependencies)
	} else if subcommand == "new" {
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
	} else if subcommand == "install" {
		reader := bufio.NewReader(os.Stdin)
		location := prompt("Enter the github url: ", reader)
		version := prompt("Enter a version code: ", reader)

		data := environment.ModuleDescription{Location: location, Version: version}
		downloader.Download(data)
		fmt.Println("Installed")
	}
}

// ref https://tutorialedge.net/golang/reading-console-input-golang/
func prompt(question string, reader *bufio.Reader) string {
	fmt.Print(question)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\r\n", "", -1)
}

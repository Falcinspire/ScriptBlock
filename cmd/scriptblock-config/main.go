package main

import (
	"fmt"
	"os"

	"github.com/falcinspire/scriptblock/environment"
	"github.com/falcinspire/scriptblock/home"
)

func main() {
	location := os.Args[1]
	version := os.Args[2]
	data := home.GetModuleConfig(environment.ModuleDescription{
		Location: location,
		Version:  version,
	})
	fmt.Println(data.Name)
	fmt.Println(data.Description)
	fmt.Println(data.Author)
	fmt.Println(data.Dependencies)
}

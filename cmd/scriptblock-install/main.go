package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/falcinspire/scriptblock/downloader"
	"github.com/falcinspire/scriptblock/environment"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	location := prompt("Enter the github url: ", reader)
	version := prompt("Enter a version code: ", reader)

	data := environment.ModuleDescription{Location: location, Version: version}
	downloader.Download(data)
	fmt.Println("Installed")
}

// ref https://tutorialedge.net/golang/reading-console-input-golang/
// duplicated code from another cmd package
func prompt(question string, reader *bufio.Reader) string {
	fmt.Print(question)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\r\n", "", -1)
}

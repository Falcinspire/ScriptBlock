package main

import (
	"os"

	"github.com/falcinspire/scriptblock/maintool"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	module := os.Args[1]
	output := os.Args[2]
	maintool.DoModule(module, output)
}

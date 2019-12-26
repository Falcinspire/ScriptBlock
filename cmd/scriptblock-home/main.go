package main

import (
	"fmt"

	"github.com/falcinspire/scriptblock/environment"
)

func main() {
	fmt.Println(environment.GetHomePath())
}

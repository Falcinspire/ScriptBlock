package downloader

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/falcinspire/scriptblock/environment"
	"github.com/falcinspire/scriptblock/home"
)

func Download(data environment.ModuleDescription) { //TODO maybe make ModuleDescription a pointer
	home.MakeModulePath(data)
	fmt.Println("Cloning " + "https://" + data.Location)
	cmd := exec.Command("git", "clone", "https://"+data.Location, data.Version) //TODO maybe include this in input? or filter it out?
	cmd.Dir = environment.GetAbreviatedModulePath(data)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// TODO checkout version
	gitHome := filepath.Join(environment.GetModulePath(data), ".git")
	fmt.Println("Deleting " + gitHome)
	os.RemoveAll(gitHome)
}

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

// BUG: works first time, but cannot be run twice on same input without failing
// git clone probably fails bc folder is not empty
func Download(data environment.ModuleDescription) { //TODO maybe make ModuleDescription a pointer
	home.MakeModulePath(data)
	fmt.Println("Cloning " + "https://" + data.Location)
	cmd := exec.Command("git", "clone", "https://"+data.Location, data.Version) //TODO maybe include this in input? or filter it out?
	cmd.Dir = environment.GetAbreviatedModulePath(data)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if data.Version != "latest" {
		versionCmd := exec.Command("git", "checkout", "tags/"+data.Version)
		versionCmd.Dir = environment.GetModulePath(data)
		versionErr := versionCmd.Run()
		if versionErr != nil {
			log.Fatal(err)
		}
	}
	gitHome := filepath.Join(environment.GetModulePath(data), ".git")
	fmt.Println("Deleting " + gitHome)
	os.RemoveAll(gitHome)
}

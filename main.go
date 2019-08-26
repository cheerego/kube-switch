package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/manifoldco/promptui"
)

var kubeDir string

func init() {
	var dir string
	var err error
	dir, err = homedir.Dir()

	if err != nil {
		color.Error.Println(err)
		panic(err)
	}
	kubeDir = path.Join(dir, ".kube")

}

func main() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: GetFiles(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	cp(fmt.Sprintf("cp %s %s", path.Join(kubeDir, result), path.Join(kubeDir, "config")))
}

func GetFiles() []string {
	var fileInfos []os.FileInfo
	var files []string
	var err error

	fileInfos, err = ioutil.ReadDir(kubeDir)
	if err != nil {
		color.Error.Println(err)
	}

	for _, file := range fileInfos {
		if !file.IsDir() && file.Name() != "config" {
			files = append(files, file.Name())
		}
	}

	if err != nil {
		color.Error.Println(err)
	}
	return files
}

func cp(cmd string) {
	command := exec.Command("bash", "-c", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	e := command.Run()
	if e != nil {
		panic(e)
	}
}

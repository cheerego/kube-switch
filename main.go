package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/gcli/v2/interact"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp(func(app *gcli.App) {
		app.Version = "1.0.6"
		app.Description = "this is a kube config switch cli applicaton"
		app.DefaultCommand("switch")
	})
	app.Add(&gcli.Command{
		Name:    "switch",
		UseFor:  "this is a description <info>message</> for command",
		Aliases: []string{"sw"},
		Func: func(cmd *gcli.Command, args []string) error {
			var fileInfos []os.FileInfo
			var files []string
			var err error
			var dir string

			dir, err = homedir.Dir()

			if err != nil {
				color.Error.Println(err)
			}

			kubeDir := path.Join(dir, ".kube")

			fileInfos, err = ioutil.ReadDir(kubeDir)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range fileInfos {
				if !file.IsDir() && file.Name() != "config" {
					files = append(files, file.Name())
				}
			}

			if err != nil {
				color.Error.Println(err)
			}
			in := interact.SingleSelect("Your Kube config", files, "")
			int, _ := strconv.Atoi(in)

			cp(fmt.Sprintf("cp %s %s", path.Join(kubeDir, files[int]), path.Join(kubeDir, "config")))
			return nil
		},
	})

	app.Run()
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

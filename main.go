package main

import (
	"github.com/urfave/cli"
	"github.com/scf0920/scf-docker/command"
)

func main(){
	app := cli.NewApp()
	app.Name = "scf-docker"
	app.Usage = "This is a simple container"
	app.Commands = []cli.Command{
		command.initCommand,
		command.runCommand,
	}
}

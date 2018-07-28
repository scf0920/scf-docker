package main

import (
	"github.com/urfave/cli"
	"github.com/scf0920/scf-docker/command"
	"log"
)

func main(){
	app := cli.NewApp()
	app.Name = "scf-docker"
	app.Usage = "This is a simple container"
	app.Commands = []cli.Command{
		command.initCommand,
		command.runCommand,
	}
	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

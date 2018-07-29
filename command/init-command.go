package command

import (
	"github.com/urfave/cli"
	"log"
)

var InitCommand = cli.Command{
	Name: "init",
	Usage: "Like docker deamon. Init a container process",
	Action: func(c *cli.Context) error {
		commands := c.Args().Get(0)
		log.Println(commands)
		return nil
	},
}
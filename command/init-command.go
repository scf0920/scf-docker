package command

import (
	"github.com/urfave/cli"
	"log"
)

var initCommand = cli.Command{
	Name: "init",
	Usage: "Like docker deamon. Init a container process",
	Action: func(c *cli.Context) error {
		commands := c.Args().Get(0)
		log.Print(commands)
		return nil
	},
}
package command

import (
	"github.com/urfave/cli"
)

var initCommand = cli.Command{
	Name: "init",
	Usage: "Like docker deamon. Init a container process",
	Action: func(c *cli.Context) error {
		return nil
	},
}
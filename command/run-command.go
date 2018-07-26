package command

import (
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: "Like docker client. Create a container",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			log.Error("Missing run command arguments.")
			return nil
		}
		
		return nil
	},
}
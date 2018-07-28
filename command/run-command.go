package command

import (
	"log"
	"github.com/urfave/cli"
	"os/exec"
	"syscall"
	"os"
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
			log.Panicln("Missing run command arguments.")
			return nil
		}
		tty := c.Bool("it") || c.Bool("ti")
		commands := c.Args().Get(0)
		args := []string{"init", commands}
		cmd := exec.Command("/proc/self/exe", args...)
        cmd.SysProcAttr = &syscall.SysProcAttr{
        	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
		    syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
        }
	    if tty {
		    cmd.Stdin = os.Stdin
		    cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
		
		if err := cmd.Start(); err != nil {
			log.Println(err)
		}
		cmd.Wait()
		os.Exit(-1)

		return nil
	},
}
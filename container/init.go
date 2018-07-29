package container

import (
	"log"
	"syscall"
	"os"
)

func RunContainerInitProcess(command string, args []string) error {
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		log.Panicln(err.Error())
	}
	return nil
}
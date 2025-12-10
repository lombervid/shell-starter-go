package commands

import (
	"fmt"
	"os"
)

type cd struct{}

func (c cd) Exec(args ...string) {
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "cd: too many arguments")
		return
	}

	path := args[0]
	ok := c.chdir(path)

	if ok == false {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
	}
}

func (c *cd) chdir(path string) bool {
	if path == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return false
		}
		path = homeDir
	}

	err := os.Chdir(path)
	if err != nil {
		return false
	}

	return true
}

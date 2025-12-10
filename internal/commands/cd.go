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
	if path == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd: : No such file or directory")
		}
		path = homeDir
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
	}
}

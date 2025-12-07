package commands

import (
	"fmt"
	"os"
)

type typeCmd struct{}

func (c typeCmd) Exec(args ...string) {
	for _, arg := range args {
		_, ok := Builtin[arg]
		if ok {
			fmt.Println(arg, "is a shell builtin")
			continue
		}

		if exec, isExecutable := FindExecutable(arg); isExecutable {
			fmt.Println(arg, "is", exec.path)
			continue
		}

		fmt.Fprintln(os.Stderr, arg+": not found")
	}
}

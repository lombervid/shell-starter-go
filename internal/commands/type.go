package commands

import (
	"fmt"
	"os"
)

type typeCmd struct{}

func (c typeCmd) Exec(args string) {
	_, ok := Builtin[args]
	if ok {
		fmt.Println(args, "is a shell builtin")
		return
	}

	if exec, isExecutable := findExecutable(args); isExecutable {
		fmt.Println(args, "is", exec.path)
		return
	}

	fmt.Fprintln(os.Stderr, args+": not found")
}

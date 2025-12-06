package commands

import (
	"fmt"
	"os"
)

type typeCmd struct{}

func (c typeCmd) Exec(args string) {
	if _, ok := Builtin[args]; !ok {
		fmt.Fprintln(os.Stderr, args+": not found")
		return
	}
	fmt.Println(args, "is a shell builtin")
}

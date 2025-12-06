package builtin

import "fmt"

type typeCmd struct{}

func (c typeCmd) Execute(args string) {
	if _, ok := Commands[args]; !ok {
		fmt.Println(args + ": not found")
		return
	}
	fmt.Println(args, "is a shell builtin")
}

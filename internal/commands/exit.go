package commands

import "os"

type exit struct{}

func (c exit) Execute(args string) {
	os.Exit(0)
}

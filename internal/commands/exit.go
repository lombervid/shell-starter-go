package commands

import "os"

type exit struct{}

func (c exit) Exec(args string) {
	os.Exit(0)
}

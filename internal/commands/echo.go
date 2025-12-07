package commands

import (
	"fmt"
)

type echo struct{}

func (c echo) Exec(args ...string) {
	var argAny = make([]any, len(args))
	for i, v := range args {
		argAny[i] = v
	}
	fmt.Println(argAny...)
}

package commands

import (
	"fmt"
)

type echo struct{}

func (c echo) Exec(args string) {
	fmt.Println(args)
}

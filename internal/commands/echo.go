package commands

import (
	"fmt"
)

type echo struct{}

func (c echo) Execute(args string) {
	fmt.Println(args)
}

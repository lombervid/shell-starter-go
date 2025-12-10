package commands

import (
	"fmt"
	"os"
)

type pwd struct{}

func (c pwd) Exec(args ...string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Println(path)
}

package main

import (
	"github.com/lombervid/shell-starter-go/internal"
)

func main() {
	for {
		input := internal.ReadInput()
		internal.HandleInput(input)
	}
}

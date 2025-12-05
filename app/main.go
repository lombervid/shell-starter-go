package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	fmt.Print("$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	command = command[:len(command)-1]
	switch command {
	case "exit":
		os.Exit(0)
	default:
		fmt.Println(command + ": command not found")
	}
}

func main() {
	for {
		repl()
	}
}

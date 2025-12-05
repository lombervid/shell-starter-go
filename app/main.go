package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	fmt.Print("$ ")

	inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	inputString = strings.TrimSpace(inputString)
	command, cmdArgs, _ := strings.Cut(inputString, " ")
	cmdArgs = strings.TrimSpace(cmdArgs)

	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(cmdArgs)
	default:
		fmt.Println(command + ": command not found")
	}
}

func main() {
	for {
		repl()
	}
}

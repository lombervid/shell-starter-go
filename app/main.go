package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lombervid/shell-starter-go/internal/commands"
)

func printPrompt() {
	fmt.Print("$ ")
}

func readInput() string {
	printPrompt()

	inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(inputString)
}

func parseCommand(input string) (cmd, args string) {
	cmd, args, _ = strings.Cut(input, " ")
	cmd = strings.TrimSpace(cmd)
	args = strings.TrimSpace(args)
	return
}

func main() {
	for {
		cmd, args := parseCommand(readInput())
		command, ok := commands.Builtin[cmd]

		if !ok {
			fmt.Fprintln(os.Stderr, cmd+": command not found")
			continue
		}

		command.Execute(args)
	}
}

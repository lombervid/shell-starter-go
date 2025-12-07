package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/lombervid/shell-starter-go/internal/commands"
)

func printPrompt() {
	fmt.Print("$ ")
}

func ReadInput() string {
	printPrompt()

	inputString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Convert multiple space characters into a single space
	inputString = regexp.MustCompile("\\s+").ReplaceAllString(inputString, " ")
	return strings.TrimSpace(inputString)
}

func parseCommand(input string) (cmd string, args []string) {
	cmd, argString, _ := strings.Cut(input, " ")
	cmd = strings.TrimSpace(cmd)
	args = strings.Split(argString, " ")
	return
}

func HandleInput(input string) {
	cmd, args := parseCommand(input)

	if command, ok := commands.Builtin[cmd]; ok {
		command.Exec(args...)
		return
	}

	if command, ok := commands.FindExecutable(cmd); ok {
		command.Exec(args...)
		return
	}

	fmt.Fprintln(os.Stderr, cmd+": command not found")
}

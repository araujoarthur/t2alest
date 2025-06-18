package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func REPLStartLoop() {
	commands := make(CommandList)

	commands.RegisterCommand("ping", Command{"type 'ping' and wait for the answer", "this is a test command", func(args ...string) error {
		fmt.Println("pong")
		return nil
	}})

	commands.RegisterCommand("exit", Command{"no flags are available for this command", "immediately exits the application", func(args ...string) error {
		os.Exit(0)
		return nil
	}})

	fmt.Println("Welcome to T2Alest (R)ead-(E)val-(P)rint (L)oop")
	fmt.Println("Remember: All paths are presumed to be relative to root (./)")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := sanitizer(scanner.Text())

		commands[input[0]].Callback(input[1:]...)
	}
}

func sanitizer(t string) []string {
	return strings.Fields(strings.ToLower(t))
}

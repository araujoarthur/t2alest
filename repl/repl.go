package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/araujoarthur/t2alest/tree"
)

func REPLStartLoop() {
	commands := GetCommands()

	fmt.Println("Welcome to T2Alest (R)ead-(E)val-(P)rint (L)oop")
	fmt.Println("Remember: All paths are presumed to be relative to root (./)")

	t := tree.CreateTree()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := sanitizer(scanner.Text())
		command, ok := commands[input[0]]
		if !ok {
			if input[0] == "q" {
				break
			}
			fmt.Printf("Command '%s' does not exist.\n", input[0])
			continue
		}

		command.Callback(t, input[1:]...)
	}
}

func sanitizer(t string) []string {
	return strings.Fields(strings.ToLower(t))
}

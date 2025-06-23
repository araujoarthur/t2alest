package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/araujoarthur/t2alest/tree"
	"github.com/google/shlex"
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

		err := command.Callback(t, input[1:]...)
		if err != nil {
			fmt.Printf("An error happened: \n%s\n", err)
		}
	}
}

func sanitizer(t string) []string {
	separated, err := shlex.Split(strings.ToLower(t))
	if err != nil {
		panic("should have no errors in shlex!")
	}

	return separated
}

func contains(slice []string, val string) (bool, int) {
	for idx, it := range slice {
		if it == val {
			return true, idx
		}
	}
	return false, 0
}

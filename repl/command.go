package repl

import (
	"fmt"
	"os"

	"github.com/araujoarthur/t2alest/tree"
)

type CommandCallback func(*tree.Tree, ...string) error
type CommandList map[string]Command

type Command struct {
	Usage    string
	HelpText string
	Callback CommandCallback
}

func (cl CommandList) registerCommand(name string, command Command) {
	cl[name] = command
}

func GetCommands() CommandList {
	newCl := make(CommandList)

	newCl.registerCommand("ping", Command{"type 'ping' and wait for the answer", "this is a test command", func(t *tree.Tree, args ...string) error {
		fmt.Println("pong")
		return nil
	}})

	newCl.registerCommand("exit", Command{"no flags are available for this command", "immediately exits the application", func(t *tree.Tree, args ...string) error {
		os.Exit(0)
		return nil
	}})

	newCl.registerCommand("ls", Command{"ls ['PATH']", "lists the content of a directory. If no path is given, it will list the contents of the current directory", func(t *tree.Tree, args ...string) error {
		fmt.Println("Não implementado")
		return nil
	}})

	newCl.registerCommand("mkdir", Command{"mkdir [-r] 'PATH'", "creates a directory, if the -r flag is present it will create all folders that does not exist in the given path", func(t *tree.Tree, args ...string) error {
		fmt.Println("Não implementado")
		return nil
	}})

	newCl.registerCommand("rm", Command{"rm [-r] 'PATH'", "removes a directory or file in PATH, if PATH is a directory and contains children the command will fail unless the -r flag is present", func(t *tree.Tree, args ...string) error {
		fmt.Println("Não implementado")
		return nil
	}})

	newCl.registerCommand("touch", Command{"touch 'PATH'", "creates an empty file at PATH. If any of the directories in path does not exist this command fails", func(t *tree.Tree, args ...string) error {
		fmt.Println("Não implementado")
		return nil
	}})

	newCl.registerCommand("find", Command{"find [-s] 'NAME'", "looks for a file or directory by 'NAME'. If the flag -s is not set, the lookup will happen in all folders and subfolders, otherwise it will perform a shallow lookup only in the root folder. ", func(t *tree.Tree, args ...string) error {
		fmt.Println("Não implementado")
		return nil
	}})

	newCl.registerCommand("help", Command{"no flags are available for this command", "prints help about the application commands", func(t *tree.Tree, args ...string) error {
		fmt.Printf("-- HELP --\n")
		for k, v := range newCl {
			fmt.Printf("%s \t-\t [%s] \t %s\n", k, v.Usage, v.HelpText)
		}

		return nil
	}})

	return newCl
}

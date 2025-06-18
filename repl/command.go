package repl

type CommandCallback func(...string) error
type CommandList map[string]Command

type Command struct {
	Usage    string
	HelpText string
	Callback CommandCallback
}

func (cl CommandList) RegisterCommand(name string, command Command) {
	cl[name] = command
}

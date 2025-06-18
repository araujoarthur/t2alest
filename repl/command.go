package repl

type CommandCallback func(...any) (string, error)
type CommandList map[string]Command

type Command struct {
	Usage    string
	HelpText string
	Callback CommandCallback
}

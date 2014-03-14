package commands

// CommandResult represents the result of the execution of a command
type CommandResult interface {
	GetResults() string
}

//
type Command interface {
	Execute(arg string) (CommandResult, error)
}

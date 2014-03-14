package commands

type ReverseCommand struct {
	Name string
}

type ReverseCommandResult struct {
	Result string
}

func (r *ReverseCommandResult) GetResults() string {
	return r.Result
}

func (r *ReverseCommand) Execute(arg string) (CommandResult, error) {
	runes := []rune(arg)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &ReverseCommandResult{Result: string(runes)}, nil
}

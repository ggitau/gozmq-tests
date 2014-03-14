package commands

type ReverseCommand struct{
	Name string
}
type ReverseCommandResult{
	Result string
}

func (r *ReverseCommandResult) GetResult() string {
	return r.Result
}

func (r *ReverseCommand) Execute(arg string)(CommandResult,error) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return ReverseCommand{Result:string(runes)}
}

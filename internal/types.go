package internal

type Command interface {
	Execute(args []string) error
}

type Process interface {
	ListProcesses() ([]Process, error)
	StartProcess() error
	EndProcessByPID() error
}

package cli

import (
	"fmt"
	"strings"

	"github.com/1AdrianM/winaxe/internal/process"
)

type Commander struct {
	p *process.Process
}

func NewCommand(p *process.Process) *Commander {
	return &Commander{p}
}
func (c *Commander) Execute(command string, options string) error {
	fmt.Print("Calling Execute")

	fmt.Print("Checking if theres options")
	comm := strings.ToLower(command) // convert to lowercase
	fmt.Println("Optiones:", options)
	switch comm {
	case "start":
		if err := c.p.StartProcess(options); err != nil {
			return fmt.Errorf("error start process command:%w \n", err)
		}
	case "kill":
		if err := c.p.EndProcessByPID(options); err != nil {
			return fmt.Errorf("error end processByPID command: %w \n", err)
		}
	case "--list-processes":
		if err := c.p.ListProcesses(); err != nil {
			return fmt.Errorf("error list process command:%w \n", err)
		}
	case "--help":
		fmt.Print("Common Commands:\n axe     shows the  process and system info  \n start         start the program \n kill       ends process \n list-processes     list all processes \n help       show this help")
		return nil

	}
	return nil
}

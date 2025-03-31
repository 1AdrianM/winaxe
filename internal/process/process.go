package process

import (
	"fmt"
	"os/exec"
	"strconv"
	"syscall"
)

const (
	PROCESS_TERMINATE = 0x0001
	EXIT_CODE         = 1
)

type Process struct {
	PID          int
	PName        string
	Status       ProcessStatus
	PUserName    string
	PDescription string
}
type ProcessStatus string

const (
	ProcessStatusRunning   ProcessStatus = "running"
	ProcessStatusStopped   ProcessStatus = "stopped"
	ProcessStatusUnknown   ProcessStatus = "unknown"
	ProcessStatusSuspended ProcessStatus = "suspended"
)

func New() *Process {
	return &Process{}
}
func (p *Process) ListProcesses() error {
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error during command of list processes%w", err)
	}
	fmt.Print(string(output))
	return nil
}
func (p *Process) StartProcess(PName string) error {
	cmd := exec.Command(PName)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error during start process %w", err)
	}
	ID := cmd.Process.Pid
	fmt.Printf("Proceso iniciado con ID: %d", ID)
	return nil
}
func (p *Process) EndProcessByPID(PID string) error {
	ID, err := strconv.Atoi(PID)
	if err != nil {
		return fmt.Errorf("error converting PID to int: %w", err)
	}
	handle, err := syscall.OpenProcess(PROCESS_TERMINATE, false, uint32(ID))
	if err != nil {
		return fmt.Errorf("error opening process %w", err)
	}
	if err := syscall.TerminateProcess(handle, EXIT_CODE); err != nil {
		return fmt.Errorf("error terminating process %w", err)
	}

	return nil
}

func (p *Process) EndProcessByName(PName string) {
	//
}
func (p *Process) GetProcessName() (string, error) {
	return p.PName, nil
}

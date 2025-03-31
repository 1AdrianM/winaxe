package pkg

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearScreen() {
	time.Sleep(1 * time.Second)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	} else {
		cmd = exec.Command("clear") // Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

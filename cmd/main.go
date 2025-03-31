package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	cli "github.com/1AdrianM/winaxe/internal/CLI"
	"github.com/1AdrianM/winaxe/internal/process"
	"github.com/1AdrianM/winaxe/pkg"
)

func main() {
	loading := "*"
	b := pkg.Blue()
	g := pkg.Green()
	for range 10 {
		g.Print(loading)
		time.Sleep(10 * time.Millisecond)

	}
	b.Print("\n Starting win axe \n")
	pkg.ClearScreen()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if scanner.Err() != nil {
		fmt.Printf("error:%s \n", scanner.Err())
		return
	}
	args := strings.Split(input, " ")
	var opts string
	comm := args[0]
	if len(args) == 2 {
		opts = args[1]
	}
	p := process.New()
	command := cli.NewCommand(p)
	if err := command.Execute(comm, opts); err != nil {
		fmt.Printf("Error: %s \n", err.Error())
	}
}

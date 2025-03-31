package pkg

import "github.com/fatih/color"

func Red() *color.Color {
	return color.New(color.FgRed)
}

func Blue() *color.Color {
	return color.New(color.FgBlue)

}

func Green() *color.Color {
	return color.New(color.FgGreen)

}

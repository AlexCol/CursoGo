package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Blue("Blue")
	color.Cyan("Cyan")
	color.Red("Red")

	b := color.BlueString("BlueString")
	r := color.RedString("RedString")
	fmt.Println(b, r)
}

package main

import (
	"fmt"

	"github.com/daviddengcn/go-colortext"
)

func printWithColor(str string, color ct.Color) {
	ct.ChangeColor(color, false, ct.None, false)
	fmt.Print(str)
	ct.ResetColor()
}

func printlnWithColor(str string, color ct.Color) {
	printWithColor(str+"\n", color)
}

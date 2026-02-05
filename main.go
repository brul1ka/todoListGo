package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// colors
var bold = color.New(color.Bold)
var errColor = color.New(color.BgRed, color.FgBlack)

func main() {
	if len(os.Args) < 2 {
		bold.Println("USAGE:\ntodo command...")
		return
	}

	command := os.Args[1]
	checkForCommands(command)
}

func printErrorMessage(errorText string) {
	fmt.Printf("%s %s\n", errColor.Sprint("error:"), errorText)
}

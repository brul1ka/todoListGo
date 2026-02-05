package main

import (
	"fmt"
	"os"
)

func checkForCommands(command string) {
	switch command {
	case "add":
		if len(os.Args) < 3 {
			printErrorMessage("Add task text after 'add'")
			break
		}
		text := os.Args[2]
		fmt.Println("Adding task:", text)
	default:
		printErrorMessage("Command \"" + os.Args[1] + "\" not found")
	}
}

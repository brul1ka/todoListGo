package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var taskID int
var tasks []Task

func loadTasks() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("INFO:", err)
		return
	}
	json.Unmarshal(data, &tasks)
	if len(tasks) > 0 {
		taskID = tasks[len(tasks)-1].ID + 1
	}

}

func checkForCommands(command string) {
	switch command {
	case "add":
		if len(os.Args) < 3 {
			printErrorMessage("Add task text after 'add'")
			break
		}
		text := os.Args[2]
		fmt.Printf("Adding task \"%s\"...\n", text)

		loadTasks()
		t := Task{ID: taskID, Text: text, Done: false}
		tasks = append(tasks, t)

		data, _ := json.MarshalIndent(tasks, "", "\t")
		os.WriteFile("tasks.json", data, 0644)

		fmt.Println("Added!")

	default:
		printErrorMessage("Command \"" + os.Args[1] + "\" not found")
	}
}

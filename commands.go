package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var taskID int
var tasks []Task

func loadTasks() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("INFO: no tasks.json file, creating...")
		os.WriteFile("tasks.json", []byte("[]"), 0644)
		fmt.Println("Created!")
		return
	}
	json.Unmarshal(data, &tasks)
	if len(tasks) > 0 {
		taskID = tasks[len(tasks)-1].ID + 1
	}

}

func checkForCommands(command string) {
	loadTasks()
	switch command {
	case "add":
		if len(os.Args) < 3 {
			printErrorMessage("Add task text after 'add'")
			break
		}
		text := os.Args[2]
		fmt.Printf("Adding task \"%s\"...\n", text)

		t := Task{ID: taskID, Text: text, Done: false}
		tasks = append(tasks, t)

		data, _ := json.MarshalIndent(tasks, "", "\t")
		os.WriteFile("tasks.json", data, 0644)

		fmt.Println("Added!")
	case "list":
		if len(tasks) <= 0 {
			fmt.Println("No tasks in file, nothing to show!")
		} else {
			for i := range tasks {
				bold.Printf("%d. \"%s\". Done: %v\n", tasks[i].ID, tasks[i].Text, tasks[i].Done)
			}
		}
	case "delete":
		if len(os.Args) < 3 {
			printErrorMessage("Task ID required after 'delete'")
			break
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			printErrorMessage("Invalid Task ID (must be a number)")
			break
		}

		if i := slices.IndexFunc(tasks, func(t Task) bool { return t.ID == id }); i != -1 {
			tasks = slices.Delete(tasks, i, i+1)
			data, _ := json.MarshalIndent(tasks, "", "\t")
			os.WriteFile("tasks.json", data, 0644)
			fmt.Println("Deleted!")
		} else {
			printErrorMessage("Task not found")
		}
	case "done":
		if len(os.Args) < 3 {
			printErrorMessage("Task ID required after 'done'")
			break
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			printErrorMessage("Invalid Task ID (must be a number)")
			break
		}

		if i := slices.IndexFunc(tasks, func(t Task) bool { return t.ID == id }); i != -1 {
			tasks[i].Done = true
			data, _ := json.MarshalIndent(tasks, "", "\t")
			os.WriteFile("tasks.json", data, 0644)
			fmt.Printf("Task %d marked as done!\n", id)
		} else {
			printErrorMessage("Task not found")
		}
	case "help":
		fmt.Println(helpText)

	default:
		printErrorMessage("Command \"" + os.Args[1] + "\" not found")
	}
}

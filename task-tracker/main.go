package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const taskFile = "tasks.json"

func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(taskFile)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func addTask(description string) {
	tasks, _ := loadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	now := time.Now().Format(time.RFC3339)

	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

func listTasks(filter string) {
	tasks, _ := loadTasks()
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("%d. %s (Status: %s)\n", task.ID, task.Description, task.Status)
		}
	}
}

func updateTask(id int, newDescription string) {
	tasks, _ := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveTasks(tasks)
			fmt.Printf("Task updated successfully (ID: %d)\n", id)
			return
		}
	}
}

func deleteTask(id int) {
	tasks, _ := loadTasks()
	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}
	if found {
		saveTasks(newTasks)
		fmt.Println("Task deleted.")
	} else {
		fmt.Println("Task not found.")
	}
}

func markStatus(id int, newStatus string) {
	tasks, _ := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveTasks(tasks)
			fmt.Printf("Task status updated to %s (ID: %d)\n", newStatus, id)
			return
		}
	}
	fmt.Println("Task not found.")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a command.")
		return
	}

	if args[1] == "list" {
		filter := ""
		if len(args) == 3 {
			filter = args[2]
		}
		listTasks(filter)
	}

	if args[1] == "add" && len(args) >= 3 {
		addTask(args[2])
	}

	if args[1] == "update" && len(args) >= 4 {
		id, _ := strconv.Atoi(args[2])
		updateTask(id, args[3])
	}

	if args[1] == "delete" && len(args) == 3 {
		id, _ := strconv.Atoi(args[2])
		deleteTask(id)
	}

	if args[1] == "mark-in-progress" && len(args) == 3 {
		id, _ := strconv.Atoi(args[2])
		markStatus(id, "in-progress")
	}

	if args[1] == "mark-done" && len(args) == 3 {
		id, _ := strconv.Atoi(args[2])
		markStatus(id, "done")
	}

}

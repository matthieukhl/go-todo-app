package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task

func ListTasks() {
	fmt.Println("Task list :")
	for _, task := range tasks {
		fmt.Printf("%d. %s (Completed : %t)\n", task.ID, task.Description, task.Completed)
	}
}

func ListUniqueTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			fmt.Printf("%d. %s (Completed : %t)\n", tasks[i].ID, tasks[i].Description, tasks[i].Completed)
		}
	}
}

func AddTask(description string) {
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, Description: description, Completed: false})
	fmt.Println("Task Added :", description)
}

func CompleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("Task Done :", tasks[i].Description)
			return
		}
	}
	fmt.Println("No such task with ID :", id)
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			fmt.Printf("Are you sure you want to delete task %v : %v? (Y/N)\n", task.ID, task.Description)

			var response string
			fmt.Scan(&response)

			if response == "Y" || response == "y" {
				tasks = append(tasks[:i], tasks[i+1:]...)
				fmt.Printf("Task %d has been deleted\n", id)
				return
			} else {
				fmt.Println("Task deletion aborted.")
			}
		}
	}
	fmt.Println("No such task with ID :", id)
}

func DeleteTaskNoConfirm(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

func SaveTasks() {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error while converting tasks in json :", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error while sasving tasks to file :", err)
	} else {
		fmt.Println("Tasks saved successfully.")
	}
}

func LoadTasks() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks file found.")
		} else {
			fmt.Println("Error while reading the tasks file :", err)
		}
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error while loading tasks from file :", err)
	} else {
		fmt.Println("Tasks loaded with success.")
	}
}

func GetCompletedTasks() []Task {
	var completedTasks []Task
	for _, task := range tasks {
		if task.Completed {
			completedTasks = append(completedTasks, task)
		}
	}

	if len(completedTasks) == 0 {
		fmt.Println("No tasks have been completed.")
		return nil
	}
	return completedTasks
}

func ShowCompletedTasks(completedTasks []Task) {
	for _, task := range completedTasks {
		fmt.Printf("%d. %s (Completed : %t)\n", task.ID, task.Description, task.Completed)
	}
}

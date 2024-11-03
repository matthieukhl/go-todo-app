package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matthieukhl/go-todo-app/task"
)

func main() {
	fmt.Println("Welcome to this ToDo App!")
	task.LoadTasks()

	var choice int

	for {
		showMenu()
		fmt.Print("Enter your choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			task.ListTasks()
		case 2:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter the task description : ")
			description, _ := reader.ReadString('\n')
			description = description[:len(description)-1]
			task.AddTask(description)
			task.SaveTasks()
		case 3:
			var id int
			fmt.Print("Enter the task ID to complete : ")
			fmt.Scan(&id)
			task.CompleteTask(id)
			task.SaveTasks()
		case 4:
			var id int
			fmt.Print("Enter the task ID to display : ")
			fmt.Scan(&id)
			task.ListUniqueTask(id)
		case 5:
			var id int
			fmt.Print("Enter the task ID to delete : ")
			fmt.Scan(&id)
			task.DeleteTask(id)
			task.SaveTasks()
		case 6:
			fmt.Println("Bye !")
			return
		case 7:
			task.ShowCompletedTasks(task.GetCompletedTasks())
		default:
			fmt.Println("\nNon valid choice. Please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\nMenu :")
	fmt.Println("1. Display all tasks")
	fmt.Println("2. Add task")
	fmt.Println("3. Complete task")
	fmt.Println("4. Display a specific task")
	fmt.Println("5. Delete task")
	fmt.Println("6. Quit")
	fmt.Println("7. Show completed tasks")
}

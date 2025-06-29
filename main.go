package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Creating a template for tasks
type Task struct {
	Id    int
	Title string
	Done  bool
}

var tasks []Task

func main() {
	// fmt.Println("Welcome to the TODOlist App")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nOptions: add, list, done, delete, quit")
		fmt.Print("> ")
		scanner.Scan()
		cmd := strings.TrimSpace(scanner.Text())

		switch cmd {
		case "add":
			fmt.Printf("\nEnter the title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			AddTask(title)

		case "list":
			ListTasks()

		case "done":
			fmt.Print("Enter ID of the task which you want to mark done : ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Task ID doesn't exist")
				break
			}
			MarkDone(id)

		case "delete":
			fmt.Print("Enter ID of the task which you want to delete : ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Task ID not found!\n")
				break
			}
			DeleteTask(id)

		case "quit":
			fmt.Println("Thanks!")
			return

		default:
			fmt.Print("Unknown command")
		}
	}

}

func AddTask(title string) {
	id := len(tasks) + 1
	t := Task{
		Id:    id,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, t)
	fmt.Printf("Added : %s\n", title)
}

func ListTasks() {
	fmt.Printf("\nYour Tasks:\n")
	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("%v. %v %v\n", t.Id, t.Title, status)
	}
}

func MarkDone(id int) {
	for i, t := range tasks {
		if t.Id == id {
			if t.Done {
				fmt.Printf("Task is already marked as Done ✅\n")
				return
			}
			tasks[i].Done = true
			fmt.Printf("Marked task as done\n")
			return
		}
	}
}

func DeleteTask(id int) {
	index := -1
	for i, t := range tasks {
		if t.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Print("Task doesn't exist")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Printf("Task Deleted\n")
}

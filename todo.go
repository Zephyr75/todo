package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "path/filepath"
	"time"

	"todo/menu"
	"todo/task"

	"github.com/rivo/tview"
)

func main() {

	tasks := make([]task.Task, 0)
	tasks = loadTasks()

	app := tview.NewApplication()
	list := tview.NewList()
	list.AddItem("Add", "Add a new task", 'a', func() {
		app.SetRoot(menu.AddMenu(app, list, &tasks), true)
	}).
		AddItem("List", "List all tasks", 'l', func() {
			app.SetRoot(menu.ListMenu(app, list, &tasks), true)
		}).
		AddItem("Quit", "Exit program", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}

	for _, t := range tasks {
		fmt.Println(t)
	}

	saveTasks(tasks)
}

func loadTasks() []task.Task {
	// absPath, _ := filepath.Abs("~/Documents/todos.txt")

	executablePath, err := os.Executable()
	absPath := filepath.Join(filepath.Dir(executablePath), "../todo/todos.txt")

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	tasks := make([]task.Task, 0)
	currentTask := task.Task{}
	i := 0
	for scanner.Scan() {
		switch i % 4 {
		case 0:
			currentTask.Title = scanner.Text()
		case 1:
			fmt.Sscanf(scanner.Text(), "%t %d %s", &currentTask.Done, &currentTask.Importance, &currentTask.Project)
		case 2:
			var dateStr, hourStr string
			fmt.Sscanf(scanner.Text(), "%s %s", &dateStr, &hourStr)
			currentTask.Date, _ = time.Parse("2006-01-02 15:04", dateStr+" "+hourStr)
		default:
			tasks = append(tasks, currentTask)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tasks
}

func saveTasks(tasks []task.Task) {
	// absPath, _ := filepath.Abs("/Documents/todos.txt")

	executablePath, err := os.Executable()
	absPath := filepath.Join(filepath.Dir(executablePath), "../todo/todos.txt")

	file, err := os.Create(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write to file
	for _, t := range tasks {
		if t.Done {
			continue
		}
		fmt.Fprintf(file, "%s\n", t.Title)
		fmt.Fprintf(file, "%t %d %s\n", t.Done, t.Importance, t.Project)
		fmt.Fprintf(file, "%s %s\n\n", t.Date.Format("2006-01-02"), t.Date.Format("15:04"))
	}
}

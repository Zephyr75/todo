package main

import (
	"time"
	"os"
	"bufio"
	"log"
	"fmt"

	"github.com/rivo/tview"
	"todo/menu"
	"todo/task"
)


func main() {
	tasks := make([]task.Task, 0)
	tasks = loadTasks()
	

	app := tview.NewApplication()
	list := tview.NewList()
	list.AddItem("Add TODO", "", 'a', func () {
			app.SetRoot(menu.AddMenu(app, list, &tasks), true)
		}).
		AddItem("List TODOs", "", 'l', func () {
			app.SetRoot(menu.ListMenu(app, list, &tasks), true)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}

	for _, t := range tasks {
		fmt.Println(t.Title)
		fmt.Println(t.Done)
		fmt.Println(t.Importance)
		fmt.Println(t.Project)
		fmt.Println(t.Date)
		fmt.Println()
	}
	saveTasks(tasks)
}

func loadTasks() []task.Task {
	file, err := os.Open("todos.txt")
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
					currentTask.Date, _ = time.Parse("2006-01-02 15:04", dateStr + " " + hourStr)
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
	file, err := os.Create("todos.txt")
		if err != nil {
				log.Fatal(err)
		}
		defer file.Close()

		// write to file
		for _, t := range tasks {
			fmt.Fprintf(file, "%s\n", t.Title)
			fmt.Fprintf(file, "%t %d %s\n", t.Done, t.Importance, t.Project)
			fmt.Fprintf(file, "%s %s\n\n", t.Date.Format("2006-01-02"), t.Date.Format("15:04"))
		}
}



package menu

import (
	"github.com/rivo/tview"
	"todo/task"
)

func ListMenu(app *tview.Application, list *tview.List, tasks *[]task.Task) *tview.Form {
	form := tview.NewForm()
	for i, _ := range *tasks {
		form.AddCheckbox((*tasks)[i].Title, (*tasks)[i].Done, func(checked bool) {
			(*tasks)[i].Done = checked
		})
	}
	form.AddButton("Back", func() {
		app.SetRoot(list, true)
	})
	return form
}

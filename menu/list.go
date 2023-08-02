package menu

import (
	"github.com/rivo/tview"
	"todo/task"

	"github.com/gdamore/tcell/v2"
)

func ListMenu(app *tview.Application, list *tview.List, tasks *[]task.Task) *tview.Form {
	form := tview.NewForm()
	for i := range *tasks {
		form.AddCheckbox((*tasks)[i].Title, (*tasks)[i].Done, func(checked bool) {
			(*tasks)[i].Done = checked
		})
	}
	form.AddButton("Back", func() {
		app.SetRoot(list, true)
	})
	form.SetFieldBackgroundColor(tcell.ColorDarkSlateBlue)
	form.SetButtonBackgroundColor(tcell.ColorDarkSlateBlue)
	return form
}

package menu

import (
	// "fmt"

	"github.com/rivo/tview"
	// "github.com/gdamore/tcell/v2"
	"time"
)

func AddMenu(app *tview.Application) *tview.Form {
	dateInput := tview.NewInputField().SetLabel("Due date").SetText(time.Now().Format("2006-01-02"))
	hourInput := tview.NewInputField().SetLabel("Due hour").SetText(time.Now().Format("15:04"))

	return tview.NewForm().
		AddInputField("Description", "", 20, nil, nil).
		AddDropDown("Importance", []string{"Chill.", "Mid", "Urgent"}, 0, nil).
		AddDropDown("Project", []string{"No project"}, 0, nil).
		AddCheckbox("Set date", false, func(checked bool) {
			if checked {
				dateInput.SetText("No date assigned")
				dateInput.SetDisabled(true)
			} else {
				dateInput.SetText(time.Now().Format("2006-01-02"))
				dateInput.SetDisabled(false)
			}
		}).
		AddFormItem(dateInput).
		AddCheckbox("Set hour", false, func(checked bool) {
			if checked {
				hourInput.SetText("No hour assigned")
				hourInput.SetDisabled(true)
			} else {
				hourInput.SetText(time.Now().Format("15:04"))
				hourInput.SetDisabled(false)
			}
		}).
		AddFormItem(hourInput).
		// AddTextArea("Address", "", 40, 0, 0, nil).
		// AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
		// AddCheckbox("Age 18+", false, nil).
		// AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Cancel", func() {
			app.Stop()
		})
}

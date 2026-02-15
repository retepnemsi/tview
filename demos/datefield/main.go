package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	dateField := tview.NewDateField().
		SetLabel("Enter a date (YYYY-MM-DD): ").
		SetPlaceholder("YYYY-MM-DD").
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(dateField, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

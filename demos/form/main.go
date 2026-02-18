// Demo code for the Form primitive.
package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/retepnemsi/tview"
)

func main() {
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorBlack,     // Screen
		ContrastBackgroundColor:     tcell.ColorAliceBlue, // Field background
		MoreContrastBackgroundColor: tcell.ColorAliceBlue, // Dropdown background
		BorderColor:                 tcell.ColorWhite,     // Border
		TitleColor:                  tcell.ColorYellow,    // Title in border
		GraphicsColor:               tcell.ColorWhite,
		PrimaryTextColor:            tcell.ColorBlack, // Text
		SecondaryTextColor:          tcell.ColorRed,   // Label
		TertiaryTextColor:           tcell.ColorYellow,
		InverseTextColor:            tcell.ColorWhite,
		ContrastSecondaryTextColor:  tcell.ColorNavy,
	}
	app := tview.NewApplication()
	form := tview.NewForm().
		AddFormItem(createDropDown()).
		//AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddDateField("Date", time.Now(), nil).
		AddTextArea("Address", "", 40, 0, 0, nil).
		AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func createDropDown() tview.FormItem {
	dropDown := tview.NewDropDown().
		SetLabel("Title").
		SetOptions([]string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, nil).
		SetCurrentOption(0).
		SetAllowEntry(true)

	return dropDown

}

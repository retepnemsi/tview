package tview

import (
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
)

type DateField struct {
	*Box
	// The text area providing the core functionality of the date field.
	textArea *TextArea

	// The screen width of the input area. This is fixed to 10 characters.
	fieldWidth int

	// An optional function which may reject the last character that was entered.
	accept func(text string, ch rune) bool

	// An optional function which is called when the input has changed.
	changed func(text string)

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, enter, or escape).
	done func(tcell.Key)

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)
}

func (df *DateField) GetFieldWidth() int {
	return 11
}

func (df *DateField) GetFieldHeight() int {
	return 1
}

func (df *DateField) GetDisabled() bool {
	return df.textArea.GetDisabled()
}

func NewDateField() *DateField {
	datefield := &DateField{
		Box:        NewBox(),
		textArea:   NewTextArea().SetWrap(false),
		fieldWidth: 11,
		accept:     inputFieldDateAcceptor,
	}
	datefield.textArea.SetChangedFunc(func() {
		if datefield.changed != nil {
			datefield.changed(datefield.textArea.GetText())
		}
	}).SetFocusFunc(func() {
		// Forward focus event to the input field.
		if datefield.Box.focus != nil {
			datefield.Box.focus()
		}
	})
	datefield.textArea.textStyle = tcell.StyleDefault.Background(Styles.ContrastBackgroundColor).Foreground(Styles.PrimaryTextColor)
	datefield.textArea.placeholderStyle = tcell.StyleDefault.Background(Styles.ContrastBackgroundColor).Foreground(Styles.ContrastSecondaryTextColor)
	return datefield
}

// SetDate sets the current text of the date field. This can be undone by the
// user. Calling this function will also trigger a "changed" event.
func (df *DateField) SetDate(date time.Time) *DateField {
	text := date.Format("2006-01-02")
	df.textArea.Replace(0, df.textArea.GetTextLength(), text)
	return df
}

// GetDate returns the Date of the date field. This returns an empty time.Time
// if the field is empty or contains an invalid date.
func (df *DateField) GetDate() time.Time {
	date, err := time.Parse("2006-01-02", df.textArea.GetText())
	if err != nil {
		return time.Time{}
	}
	return date
}

// SetLabel sets the text to be displayed before the input area.
func (df *DateField) SetLabel(label string) *DateField {
	df.textArea.SetLabel(label)
	return df
}

// GetLabel returns the text to be displayed before the input area.
func (df *DateField) GetLabel() string {
	return df.textArea.GetLabel()
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (df *DateField) SetLabelWidth(width int) *DateField {
	df.textArea.SetLabelWidth(width)
	return df
}

// SetPlaceholder sets the text to be displayed when the input text is empty.
func (df *DateField) SetPlaceholder(text string) *DateField {
	df.textArea.SetPlaceholder(text)
	return df
}

// SetLabelColor sets the text color of the label.
func (df *DateField) SetLabelColor(color tcell.Color) *DateField {
	df.textArea.SetLabelStyle(df.textArea.GetLabelStyle().Foreground(color))
	return df
}

// SetLabelStyle sets the style of the label.
func (df *DateField) SetLabelStyle(style tcell.Style) *DateField {
	df.textArea.SetLabelStyle(style)
	return df
}

// SetFieldBackgroundColor sets the background color of the input area.
func (df *DateField) SetFieldBackgroundColor(color tcell.Color) *DateField {
	df.textArea.SetTextStyle(df.textArea.GetTextStyle().Background(color))
	return df
}

// SetFieldTextColor sets the text color of the input area.
func (df *DateField) SetFieldTextColor(color tcell.Color) *DateField {
	df.textArea.SetTextStyle(df.textArea.GetTextStyle().Foreground(color))
	return df
}

// SetFieldStyle sets the style of the input area (when no placeholder is
// shown).
func (df *DateField) SetFieldStyle(style tcell.Style) *DateField {
	df.textArea.SetTextStyle(style)
	return df
}

// GetFieldStyle returns the style of the input area (when no placeholder is
// shown).
func (df *DateField) GetFieldStyle() tcell.Style {
	return df.textArea.GetTextStyle()
}

// SetPlaceholderTextColor sets the text color of placeholder text.
func (df *DateField) SetPlaceholderTextColor(color tcell.Color) *DateField {
	df.textArea.SetPlaceholderStyle(df.textArea.GetPlaceholderStyle().Foreground(color))
	return df
}

// SetPlaceholderStyle sets the style of the input area (when a placeholder is
// shown).
func (df *DateField) SetPlaceholderStyle(style tcell.Style) *DateField {
	df.textArea.SetPlaceholderStyle(style)
	return df
}

// GetPlaceholderStyle returns the style of the input area (when a placeholder
// is shown).
func (df *DateField) GetPlaceholderStyle() tcell.Style {
	return df.textArea.GetPlaceholderStyle()
}

// SetFormAttributes sets attributes shared by all form items.
func (df *DateField) SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) FormItem {
	df.textArea.SetFormAttributes(labelWidth, labelColor, bgColor, fieldTextColor, fieldBgColor)
	return df
}

// SetDisabled sets whether the item is disabled / read-only.
func (df *DateField) SetDisabled(disabled bool) FormItem {
	df.textArea.SetDisabled(disabled)
	if df.finished != nil {
		df.finished(-1)
	}
	return df
}

// SetChangedFunc sets a handler which is called whenever the text of the input
// field has changed. It receives the current text (after the change).
func (df *DateField) SetChangedFunc(handler func(text string)) *DateField {
	df.changed = handler
	return df
}

// SetDoneFunc sets a handler which is called when the user is done entering
// text. The callback function is provided with the key that was pressed, which
// is one of the following:
//
//   - KeyEnter: Done entering text.
//   - KeyEscape: Abort text input.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (df *DateField) SetDoneFunc(handler func(key tcell.Key)) *DateField {
	df.done = handler
	return df
}

// SetFinishedFunc sets a callback invoked when the user leaves this form item.
func (df *DateField) SetFinishedFunc(handler func(key tcell.Key)) FormItem {
	df.finished = handler
	return df
}

// Focus is called when this primitive receives focus.
func (df *DateField) Focus(delegate func(p Primitive)) {
	// If we're part of a form and this item is disabled, there's nothing the
	// user can do here so we're finished.
	if df.finished != nil && df.textArea.GetDisabled() {
		df.finished(-1)
		return
	}

	df.Box.Focus(delegate)
}

// HasFocus returns  this primitive has focus.
func (df *DateField) HasFocus() bool {
	return df.textArea.HasFocus() || df.Box.HasFocus()
}

// Blur is called when this primitive loses focus.
func (df *DateField) Blur() {
	df.textArea.Blur()
	df.Box.Blur()
}

// Draw draws this primitive onto the screen.
func (df *DateField) Draw(screen tcell.Screen) {
	df.Box.DrawForSubclass(screen, df)

	// Prepare
	x, y, width, height := df.GetInnerRect()
	if height < 1 || width < 1 {
		return
	}

	// Resize text area.
	labelWidth := df.textArea.GetLabelWidth()
	if labelWidth == 0 {
		labelWidth = TaggedStringWidth(df.textArea.GetLabel())
	}
	fieldWidth := df.fieldWidth
	if fieldWidth == 0 {
		fieldWidth = width - labelWidth
	}
	df.textArea.SetRect(x, y, labelWidth+fieldWidth, 1)
	df.textArea.setMinCursorPadding(fieldWidth-1, 1)

	// Draw text area.
	df.textArea.hasFocus = df.HasFocus() // Force cursor positioning.
	df.textArea.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (df *DateField) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return df.WrapInputHandler(func(key *tcell.EventKey, f func(p Primitive)) {
		df.inputHandler(key, f)
	})
}

func (df *DateField) inputHandler(event *tcell.EventKey, setFocus func(p Primitive)) {
	if df.textArea.GetDisabled() {
		return
	}

	// If we have an autocomplete list, there are certain keys we will
	// forward to it.

	// Finish up.
	finish := func(key tcell.Key) {
		if df.done != nil {
			df.done(key)
		}
		if df.finished != nil {
			df.finished(key)
		}
	}

	// Process special key events for the input field.
	switch key := event.Key(); key {
	case tcell.KeyEnter, tcell.KeyEscape, tcell.KeyTab, tcell.KeyBacktab:
		finish(key)
	case tcell.KeyCtrlV:
		if df.accept != nil && !df.accept(df.textArea.getTextBeforeCursor()+df.textArea.GetClipboardText()+df.textArea.getTextAfterCursor(), 0) {
			return
		}
		df.textArea.InputHandler()(event, setFocus)
	case tcell.KeyRune:
		if event.Modifiers()&tcell.ModAlt == 0 && df.accept != nil {
			// Check if this rune is accepted.
			r := event.Rune()
			if !unicode.IsDigit(r) && r != '-' {
				return
			}
			text := df.textArea.getTextBeforeCursor()
			if len(text) == 4 || len(text) == 7 {
				text += "-" + df.textArea.getTextAfterCursor()
				df.textArea.SetText(text, true)
			}
			if !df.accept(df.textArea.getTextBeforeCursor()+string(r)+df.textArea.getTextAfterCursor(), r) {
				return
			}
		}
		fallthrough
	default:
		// Forward other key events to the text area.
		df.textArea.InputHandler()(event, setFocus)
	}
}

// MouseHandler returns the mouse handler for this primitive.
func (df *DateField) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return df.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if df.textArea.GetDisabled() {
			return false, nil
		}

		// Is mouse event within the input field?
		x, y := event.Position()
		if !df.InRect(x, y) {
			return false, nil
		}

		// Forward mouse event to the text area.
		consumed, capture = df.textArea.MouseHandler()(action, event, setFocus)

		return
	})
}

// PasteHandler returns the handler for this primitive.
func (df *DateField) PasteHandler() func(pastedText string, setFocus func(p Primitive)) {
	return df.WrapPasteHandler(func(pastedText string, setFocus func(p Primitive)) {
		// Input field may be disabled.
		if df.textArea.GetDisabled() {
			return
		}

		// We may not accept this text.
		if df.accept != nil && !df.accept(df.textArea.getTextBeforeCursor()+pastedText+df.textArea.getTextAfterCursor(), 0) {
			return
		}

		// Forward the pasted text to the text area.
		df.textArea.PasteHandler()(pastedText, setFocus)
	})
}

func inputFieldDateAcceptor(text string, ch rune) bool {
	if len(text) > 10 {
		return false
	}
	if !unicode.IsDigit(ch) && ch != '-' {
		return false
	}
	if len(text) < 5 && !unicode.IsDigit(ch) {
		return false
	}
	if len(text) == 5 && ch != '-' {
		return false
	}
	if len(text) > 5 && len(text) < 8 && !unicode.IsDigit(ch) {
		return false
	}
	if len(text) == 6 && ch != '0' && ch != '1' {
		return false
	}
	if len(text) == 8 && ch != '-' {
		return false
	}
	if len(text) > 8 && len(text) < 11 && !unicode.IsDigit(ch) {
		return false
	}
	if len(text) == 9 && !(ch >= '0' && ch < '4') {
		return false
	}
	if len(text) == 10 {
		_, err := time.Parse("2006-01-02", text)
		if err != nil {
			return false
		}
	}
	return true
}

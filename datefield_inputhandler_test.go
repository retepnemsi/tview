package tview

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

type keyTap = struct {
	key    rune
	result string
}

func TestDateField_input_with_dashes(t *testing.T) {
	field := NewDateField()

	var test = []keyTap{
		{'2', "2"},
		{'0', "20"},
		{'2', "202"},
		{'6', "2026"},
		{'-', "2026-"},
		{'0', "2026-0"},
		{'2', "2026-02"},
		{'-', "2026-02-"},
		{'1', "2026-02-1"},
		{'4', "2026-02-14"},
	}

	runTest(t, test, field)
}

func TestDateField_input_without_dashes(t *testing.T) {
	field := NewDateField()

	var test = []keyTap{
		{'2', "2"},
		{'0', "20"},
		{'2', "202"},
		{'6', "2026"},
		{'0', "2026-0"},
		{'2', "2026-02"},
		{'1', "2026-02-1"},
		{'4', "2026-02-14"},
	}

	runTest(t, test, field)
}

func TestDateField_input_with_letters_and_dashes(t *testing.T) {
	field := NewDateField()

	var test = []keyTap{
		{'2', "2"},
		{'0', "20"},
		{'2', "202"},
		{'6', "2026"},
		{'a', "2026"},
		{'-', "2026-"},
		{'0', "2026-0"},
		{'2', "2026-02"},
		{'-', "2026-02-"},
		{'1', "2026-02-1"},
		{'4', "2026-02-14"},
	}

	runTest(t, test, field)
}

func TestDateField_input_with_dashes_in_wrong_place(t *testing.T) {
	field := NewDateField()

	var test = []keyTap{
		{'2', "2"},
		{'0', "20"},
		{'2', "202"},
		{'-', "202"},
		{'6', "2026"},
		{'0', "2026-0"},
		{'2', "2026-02"},
		{'-', "2026-02-"},
		{'1', "2026-02-1"},
		{'4', "2026-02-14"},
	}

	runTest(t, test, field)
}

func TestDateField_input_with_multiple_wrong_input(t *testing.T) {
	field := NewDateField()

	var test = []keyTap{
		{'2', "2"},
		{'0', "20"},
		{'2', "202"},
		{'-', "202"},
		{'a', "202"},
		{'6', "2026"},
		{'b', "2026"},
		{'0', "2026-0"},
		{'2', "2026-02"},
		{'-', "2026-02-"},
		{'-', "2026-02-"},
		{'1', "2026-02-1"},
		{'c', "2026-02-1"},
		{'d', "2026-02-1"},
		{'e', "2026-02-1"},
		{'4', "2026-02-14"},
	}

	runTest(t, test, field)
}

func runTest(t *testing.T, test []keyTap, field *DateField) {
	for _, testData := range test {
		key := tcell.NewEventKey(tcell.KeyRune, testData.key, tcell.ModNone)
		field.inputHandler(key, nil)
		assert.Equal(t, testData.result, field.textArea.GetText())
	}
}

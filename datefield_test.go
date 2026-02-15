package tview

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

// Happy path tests
func TestDateAcceptor_first_digit_of_year(t *testing.T) {
	result := inputFieldDateAcceptor("2", '2')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_second_digit_of_year(t *testing.T) {
	result := inputFieldDateAcceptor("20", '0')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_third_digit_of_year(t *testing.T) {
	result := inputFieldDateAcceptor("20", '2')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_fourth_digit_of_year(t *testing.T) {
	result := inputFieldDateAcceptor("2026", '6')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_first_separator(t *testing.T) {
	result := inputFieldDateAcceptor("2026-", '-')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_first_digit_of_month(t *testing.T) {
	result := inputFieldDateAcceptor("2026-0", '0')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_second_digit_of_month(t *testing.T) {
	result := inputFieldDateAcceptor("2026-02", '2')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_second_digit_of_month_2(t *testing.T) {
	result := inputFieldDateAcceptor("2026-09", '9')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_second_separator(t *testing.T) {
	result := inputFieldDateAcceptor("2026-02-", '-')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_first_digit_of_day(t *testing.T) {
	result := inputFieldDateAcceptor("2026-02-1", '1')
	assert.Equal(t, true, result)
}
func TestDateAcceptor_second_digit_of_day(t *testing.T) {
	result := inputFieldDateAcceptor("2026-02-14", '4')
	assert.Equal(t, true, result)
}

// Test every month, good days
func TestDateAcceptor(t *testing.T) {
	for i := 1; i < 13; i += 1 {
		text := fmt.Sprintf("2025-%02d-05", i)
		fmt.Println(text)
		result := inputFieldDateAcceptor(text, '5')
		assert.Equal(t, true, result, "Expected "+text+" to succeed")
	}
}

// Failing path tests
func TestDateAcceptor_first_digit_of_year_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1", 'a')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_second_digit_of_year_fails(t *testing.T) {
	result := inputFieldDateAcceptor("12", 'b')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_third_digit_of_year_fails(t *testing.T) {
	result := inputFieldDateAcceptor("123", 'c')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_fourth_digit_of_year_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1234", 'd')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_first_separator_fails(t *testing.T) {
	result := inputFieldDateAcceptor("12344", '4')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_first_digit_of_month_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1234-5", '-')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_second_digit_of_month_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1234-56", '-')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_second_separator_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1234-56-", '1')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_first_digit_of_day_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1234-56-7", '-')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_second_digit_of_day_fails(t *testing.T) {
	result := inputFieldDateAcceptor("1234-56-78", '-')
	assert.Equal(t, false, result)
}
func TestDateAcceptor_failing_date(t *testing.T) {
	result := inputFieldDateAcceptor("2026-02-29", '9')
	assert.Equal(t, false, result)
}

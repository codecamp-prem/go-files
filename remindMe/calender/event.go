package calender

// unicode/utf-8 package let us count no. of runes in a string
import (
	"errors"
	"unicode/utf8"
)

// Event : Title of the remainder
type Event struct {
	title string
	Date
}

// Title method(Getter) return title of remainder calender
func (e *Event) Title() string {
	return e.title
}

// SetTitle method(Setter) allows to set title of less than 30 letters
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid title")
	}
	e.title = title
	return nil
}

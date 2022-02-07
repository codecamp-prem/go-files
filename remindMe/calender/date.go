package calender

import "errors"

// Date : type defined struct
type Date struct {
	year  int
	month int
	day   int
}

// SetYear : set year
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth : set Month
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

// SetDay : set day
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}

/* Getter methods */

// Year returns set year
func (d *Date) Year() int {
	return d.year
}

//Month return set month
func (d *Date) Month() int {
	return d.month
}

//Day return set day
func (d *Date) Day() int {
	return d.day
}

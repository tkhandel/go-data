package godata

import "fmt"

type Duplicate struct {
	What  string
	Value string
}

func (d Duplicate) Error() string {
	return fmt.Sprintf("duplicate %s: %s", d.What, d.Value)
}

type Unknown struct {
	What  string
	Value string
}

func (u Unknown) Error() string {
	return fmt.Sprintf("unknown %s: %s", u.What, u.Value)
}

type ProcessingError struct {
	Err error
}

func (p ProcessingError) Error() string {
	return p.Err.Error()
}

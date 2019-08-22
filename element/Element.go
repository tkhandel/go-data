package element

import (
	"fmt"
	"github.com/pkg/errors"
)

type Element struct {
	value interface{}
}

func New(value interface{}) Element {
	return Element{value: value}
}

func (e Element) String() (string, error) {
	val, ok := e.value.(string)
	if !ok {
		return fmt.Sprint(e.value), nil
	}
	return val, nil
}

func (e Element) MustString() string {
	return e.value.(string)
}

func (e Element) Int() (int, error) {
	val, ok := e.value.(int)
	if !ok {
		return 0, errors.New("invalid cast to int")
	}
	return val, nil
}

func (e Element) MustInt() int {
	return e.value.(int)
}

func (e Element) Float() (float64, error) {
	val, ok := e.value.(float64)
	if !ok {
		return 0, errors.New("invalid cast to float")
	}
	return val, nil
}

func (e Element) MustFloat() float64 {
	return e.value.(float64)
}

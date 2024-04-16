package gtml

import (
	"log"
	"reflect"

	"github.com/4lxprime/gtml/elements"
)

type App struct {
	Element      elements.Element
	StateManager *StateManager
}

func NewApp() *App {
	return &App{
		StateManager: NewStateManager(),
	}
}

func (a *App) UseState(v interface{}) *State {
	s := &State{
		channel: make(chan interface{}),
		value:   nil,
	}

	a.StateManager.appendState(s)

	s.channel <- v

	return s
}

// this function will be used to give App main element
func (a *App) Use(el elements.Element) *App {
	a.Element = el

	return a
}

type statement struct {
	el elements.Element
}

// NOTE: if there is no Else or Elif after, you should
// add .Value() after
func If(
	condition bool,
) func(el elements.Element) statement {
	return func(el elements.Element) statement {
		if condition {
			return statement{
				el: el,
			}
		}

		return statement{
			el: &elements.EmptyEl{},
		}
	}
}

// NOTE: if there is no Else or Elif after, you should
// add .Value() after
func (s statement) Elif(
	condition bool,
) func(el elements.Element) statement {
	return func(el elements.Element) statement {
		switch s.el.(type) {
		case *elements.EmptyEl:
			if condition {
				return statement{
					el: el,
				}
			}
			return s

		default:
			return s
		}
	}
}

func (s statement) Else(el elements.Element) elements.Element {
	switch s.el.(type) {
	case *elements.EmptyEl:
		return el

	default:
		return s.el
	}
}

func (s statement) Value() elements.Element {
	return s.el
}

func For(
	init, reached int,
) func(fn func(int) elements.Element) *elements.SliceEl {
	sliceElement := &elements.SliceEl{}

	reverse := reached > init

	return func(fn func(int) elements.Element) *elements.SliceEl {
		if reverse {
			for i := init; i >= reached; i-- {
				sliceElement.AppendChild(fn(i))
			}
		}

		for i := init; i <= reached; i++ {
			sliceElement.AppendChild(fn(i))
		}

		return sliceElement
	}
}

func Each(
	slice interface{},
) func(fn func(int, any) elements.Element) *elements.SliceEl {
	sliceElement := &elements.SliceEl{}
	return func(fn func(int, any) elements.Element) *elements.SliceEl {
		if reflect.TypeOf(slice).Kind() != reflect.Slice {
			log.Println("argument must be a slice")
			return sliceElement
		}

		s := reflect.ValueOf(slice)
		for i := 0; i <= s.Len(); i++ {
			sliceElement.AppendChild(
				fn(i, s.Index(i).Interface()),
			)
		}

		return sliceElement
	}
}

func Each2[T any](
	slice []T,
) func(fn func(int, T) elements.Element) *elements.SliceEl {
	sliceElement := &elements.SliceEl{}
	return func(fn func(int, T) elements.Element) *elements.SliceEl {
		for i := 0; i < len(slice); i++ {
			sliceElement.AppendChild(
				fn(i, slice[i]),
			)
		}

		return sliceElement
	}
}

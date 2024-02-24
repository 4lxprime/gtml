package gtml

import "github.com/4lxprime/gtml/elements"

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

func (a *App) Use(el elements.Element) *App {
	a.Element = el

	return a
}

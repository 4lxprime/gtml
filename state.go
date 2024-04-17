package gtml

import (
	"context"
	"sync"
)

// the state represents a reactive value in the main component
//
// NOTE: id is used by the state manager in the runtime
// and when something is pushed through the channel,
// this data will be placed on value field
type State struct {
	id      int64
	started bool
	channel chan interface{}
	value   interface{}
}

func (s *State) Get() interface{} {
	if !s.started {
		return s.value
	}

	return <-s.channel
}

func (s *State) Set(v interface{}) { s.channel <- v }

type StateManager struct {
	states map[int64]*State
	mutex  sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

func NewStateManager() *StateManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &StateManager{
		states: make(map[int64]*State),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (m *StateManager) appendState(s *State) {
	var id int64 = int64(len(m.states))

	s.id = id

	m.states[id] = s

	go m.listenState(s)
}

func (m *StateManager) Start() {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, state := range m.states {
		state.started = true
	}
}

func (m *StateManager) listenState(s *State) {
	for {
		select {
		case <-m.ctx.Done():
			return

		case val, ok := <-s.channel:
			if !ok {
				m.mutex.Lock()
				delete(m.states, s.id)
				m.mutex.Unlock()
				return
			}
			s.value = val
		}
	}
}

func (m *StateManager) Stop() { m.cancel() }

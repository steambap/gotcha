package main

import (
	"sync"
)

type Action struct {
	Name    string
	Payload interface{}
}

type State struct {
	Name  string
	Id    int
	Stuff []interface{}
}

type Store struct {
	*State
	Mutex sync.Mutex
}

func (s *Store) Dispatch(action *Action) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	newState := getNewState(action, s.State)
	s.State = newState
}

func getNewState(action *Action, state *State) *State {
	switch action.Name {
	case "setName":
		newName := action.Payload.(string)
		state.Name = newName
		return state
	case "setId":
		newId := action.Payload.(int)
		state.Id = newId
		return state
	default:
		return state
	}
}

package manager

import "github.com/maxcelant/pomo-cli/cmd/state"

type StateManager struct {
	State state.State
}

func New(s state.State) StateManager {
	return StateManager{s}
}

func (sm *StateManager) UpdateState(newState state.State) {
	sm.State = newState
}

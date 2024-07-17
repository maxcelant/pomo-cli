
package manager

import "github.com/maxcelant/pomo-cli/cmd/state"

type StateManager struct {
  State state.State
  Intervals int
}

func New(s state.State, i int) *StateManager {
  return &StateManager{s, i}
}

func (sm *StateManager) UpdateState(newState state.State) {
  sm.State = newState
}

func (sm *StateManager) DecrementInterval() {
  sm.Intervals--
}

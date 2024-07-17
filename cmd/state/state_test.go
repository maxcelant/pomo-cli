package state

import (
	"testing"
)

func TestNewState(t *testing.T) {
  id := state.INIT
  literal := "active"
  symbol := "🟢"
  
  state := New(id, literal, symbol)
  
  if state.id != id {
    t.Errorf("Expected id %d, but got %d", id, state.id)
  }
  
  if state.literal != literal {
    t.Errorf("Expected literal %s, but got %s", literal, state.literal)
  }
  
  if state.symbol != symbol {
    t.Errorf("Expected symbol %s, but got %s", symbol, state.symbol)
  }
}

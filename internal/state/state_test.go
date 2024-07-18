package state

import (
	"testing"
)

func TestNewState(t *testing.T) {
	id := INIT
	literal := "active"
	symbol := "🟢"

	state := New(id, literal, symbol)

	if state.Id != id {
		t.Errorf("Expected id %d, but got %d", id, state.Id)
	}

	if state.Literal != literal {
		t.Errorf("Expected literal %s, but got %s", literal, state.Literal)
	}

	if state.Symbol != symbol {
		t.Errorf("Expected symbol %s, but got %s", symbol, state.Symbol)
	}
}

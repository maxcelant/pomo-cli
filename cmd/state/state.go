
package state

type ID int

const (
  INIT ID = iota
  ACTIVE
  BREAK
  WAITING
  PAUSE
  DONE
)

type State struct {
  Id ID 
  Literal string
  Symbol string
}

func New(id ID, literal string, symbol string) State {
  return State{
    Id:id,
    Literal: literal,
    Symbol: symbol,
  }
}

var States map[ID]State

func init() {
	States = map[ID]State{
		INIT:    New(INIT, "Initial", "🔵"),
		ACTIVE:  New(ACTIVE, "Active", "🟢"),
		BREAK:   New(BREAK, "Break", "🔴"),
		WAITING: New(WAITING, "Waiting", "⏱️"),
		PAUSE:   New(PAUSE, "Pause", "⏸️"),
		DONE:    New(DONE, "Done", "✅"),
	}
}

func Get(id ID) State {
	return States[id]
}

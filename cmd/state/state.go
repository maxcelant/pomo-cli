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
	Id      ID
	Literal string
	Symbol  string
  Duration int 
}

func New(id ID, literal string, symbol string, duration int) State {
	return State{
		Id:      id,
		Literal: literal,
		Symbol:  symbol,
    Duration: duration,
	}
}

var States = map[ID]State{
	INIT:    New(INIT, "Initial", "🔵", 0),
	ACTIVE:  New(ACTIVE, "Active", "🟢", 25),
	BREAK:   New(BREAK, "Break", "🔴", 10),
	WAITING: New(WAITING, "Waiting", "🕒", 0),
	PAUSE:   New(PAUSE, "Pause", "⏸️", 0),
	DONE:    New(DONE, "Done", "✅", 0),
}

func Get(id ID) State {
	return States[id]
}

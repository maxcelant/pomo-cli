package state

import (
	"log"

	"github.com/maxcelant/pomo-cli/internal/fileio"
)

type ID int

const (
	INIT ID = iota
	ACTIVE
	REST
	WAITING
	PAUSE
	DONE
)

type State struct {
	Id       ID
	Literal  string
	Symbol   string
	Duration int
}

func New(id ID, literal string, symbol string, duration int) State {
	return State{
		Id:       id,
		Literal:  literal,
		Symbol:   symbol,
		Duration: duration,
	}
}

var States map[ID]State

func init() {
	activeDuration := 1500 // 25 mins
	restDuration := 600    // 10 mins

	if config, err := fileio.ReadFromLocalYaml("pomo.yaml"); err != nil {
		log.Printf("Error reading config file: %s. Using default durations.", err)
	} else {
		activeDuration = config.Pomo.Active * 60
		restDuration = config.Pomo.Rest * 60
	}

	States = map[ID]State{
		INIT:    New(INIT, "Initial", "🔵", 0),
		ACTIVE:  New(ACTIVE, "Active", "🟢", activeDuration),
		REST:    New(REST, "Rest", "🔴", restDuration),
		WAITING: New(WAITING, "Waiting", "🕒", 0),
		PAUSE:   New(PAUSE, "Pause", "⏸️", 0),
		DONE:    New(DONE, "Done", "✅", 0),
	}
}

func Get(id ID) State {
	return States[id]
}

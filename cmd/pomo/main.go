package main

import (
	"fmt"
	"os"

	"github.com/maxcelant/pomo-cli/cmd/manager"
	"github.com/maxcelant/pomo-cli/cmd/state"
	"github.com/maxcelant/pomo-cli/cmd/timer"
	"github.com/maxcelant/pomo-cli/cmd/screen"
	"github.com/maxcelant/pomo-cli/cmd/session"
)


func handleStartCommand(session *session.Session, subcommands []string) {
	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	for {
    session.Loop(state.ACTIVE)
    session.Loop(state.BREAK)
    session.IncrementInterval()
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		screen.Usage()
		os.Exit(1)
	}

  sm := manager.New(state.Get(state.INIT))
  timer := timer.New()
  session := session.New(sm, timer, 0)

	switch args[1] {
	case "start":
		handleStartCommand(session, args[2:])
	default:
		screen.Usage()
	}
}

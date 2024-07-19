package main

import (
	"fmt"
	"os"

	"github.com/maxcelant/pomo-cli/internal/manager"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"github.com/maxcelant/pomo-cli/internal/session"
	"github.com/maxcelant/pomo-cli/internal/state"
	"github.com/maxcelant/pomo-cli/internal/subcommand"
	"github.com/maxcelant/pomo-cli/internal/timer"
)

func handleStartCommand(session *session.Session, subcommands []string) {
	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	for {
		session.Loop(state.ACTIVE)
		session.Loop(state.REST)
		session.IncrementInterval()
	}
}

func handleConfigCommand(subcommands []string) {
	out, err := subcommand.Handler(subcommands, map[string]interface{}{"active": 0, "rest": 0})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Println("Subcommand values:", out)
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
	case "config":
		handleConfigCommand(args[2:])
	default:
		screen.Usage()
	}
}

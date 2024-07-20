package main

import (
	"fmt"
	"os"

	"github.com/maxcelant/pomo-cli/internal/manager"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"github.com/maxcelant/pomo-cli/internal/session"
	"github.com/maxcelant/pomo-cli/internal/state"
	"github.com/maxcelant/pomo-cli/internal/timer"
	"github.com/maxcelant/pomo-cli/internal/command"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		screen.Usage()
		os.Exit(1)
	}

	sm := manager.New(state.Get(state.INIT))
	timer := timer.New()
	session := session.New(sm, timer, 1)

  handler, err := command.NewHandler(args[1], session, args[2:])
  if err != nil {
    fmt.Println("Error: ", err)
    screen.Usage()
    os.Exit(1)
  }

  handler.Handle()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/maxcelant/pomo-cli/cmd/manager"
	"github.com/maxcelant/pomo-cli/cmd/state"
	"github.com/maxcelant/pomo-cli/cmd/timer"
	"github.com/maxcelant/pomo-cli/cmd/screen"
)

func awaitUserRes(sm *manager.StateManager) {
  screen.Clear()
  if sm.State.Id == state.ACTIVE {
    fmt.Printf("🍎 Active session done! Ready to start break?")
  } else {
    fmt.Printf("🍎 Break session done! Ready to start studying?")
  }
  fmt.Printf("   [Enter] to continue, or [Q]uit: ")
	reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)
  if input == "q" || input == "Q" {
    fmt.Printf("  Exiting gracefully...")
    os.Exit(0)
  }
}

func handleStartCommand(sm *manager.StateManager, timer *timer.Timer, subcommands []string) {
	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	for {
    sm.UpdateState(state.Get(state.ACTIVE))
    timer.SetDuration(sm.State.Duration)
		timer.Time(sm)
    awaitUserRes(sm)
    sm.UpdateState(state.Get(state.BREAK))
    timer.SetDuration(sm.State.Duration)
		timer.Time(sm)
    awaitUserRes(sm)
    sm.IncrementInterval()
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		screen.Usage()
		os.Exit(1)
	}

  sm := manager.New(state.Get(state.INIT), 0)
  timer := timer.New()

	switch args[1] {
	case "start":
		handleStartCommand(sm, timer, args[2:])
	default:
		screen.Usage()
	}
}

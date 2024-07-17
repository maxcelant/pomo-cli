package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/maxcelant/pomo-cli/cmd/manager"
	"github.com/maxcelant/pomo-cli/cmd/state"
)

func usage() {
	fmt.Printf("Usage: pomo [command]\n\n")
  fmt.Println("       Commands           Actions")
  fmt.Println("       start              Begin a simple pomodoro interval session.")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func timer(sm *manager.StateManager) {
	for t := 0; t < sm.State.Duration; t++ {
		select {
		case <-time.After(1 * time.Second):
      fmt.Printf("t: %d", t)
			clearScreen()
			fmt.Println("🍎 Time to focus")
			fmt.Printf("   State: %s %s\n", sm.State.Literal, sm.State.Symbol)
			fmt.Printf("   Interval: %d\n", sm.Intervals)
			fmt.Printf("   Time Remaining: %ds\n", sm.State.Duration-t)
		}
	}
}

func awaitUserRes(sm *manager.StateManager) {
  clearScreen()
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

func handleStartCommand(sm *manager.StateManager, subcommands []string) {
	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	for ; sm.Intervals > 0; sm.DecrementInterval() {
    sm.UpdateState(state.Get(state.ACTIVE))
		timer(sm)
    awaitUserRes(sm)
    sm.UpdateState(state.Get(state.BREAK))
		timer(sm)
    awaitUserRes(sm)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

  sm := manager.New(state.Get(state.INIT), 3)

	switch args[1] {
	case "start":
		handleStartCommand(sm, args[2:])
	default:
		usage()
	}
}

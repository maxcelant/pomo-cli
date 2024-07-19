package main

import (
	"fmt"
	"os"
  "strconv"

	"github.com/maxcelant/pomo-cli/internal/manager"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"github.com/maxcelant/pomo-cli/internal/session"
	"github.com/maxcelant/pomo-cli/internal/state"
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

func handleConfigCommand(session *session.Session, subcommands []string) {

  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Error: input was not given for one of your flags")
      os.Exit(1)
    }
  }()

  flags := map[string]string{"-a": "active", "--active": "active", "-r": "rest", "--rest": "rest"}
  states := map[string]int{"active": 0, "rest": 0}

  for i := 0; i < len(subcommands); i++ {
    flag := subcommands[i]

    f, found := flags[flag]
    if !found {
      fmt.Printf("flag '%s' is not a viable flag\n", flag)
      return
    }

    if i+1 > len(subcommands) {
      fmt.Printf("flag '%s' expects a value but none was provided\n", flag)
      return
    }

    nextValue := subcommands[i+1]
    duration, err := strconv.Atoi(nextValue)
    if err != nil {
      fmt.Printf("value for flag '%s' is not a valid integer: %s\n", flag, nextValue)
      return
    }

    states[f] = duration
    i++ 
  }

  fmt.Println("Subcommand values:", states)
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
    handleConfigCommand(session, args[2:])
	default:
		screen.Usage()
	}
}

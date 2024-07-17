package main

import (
	"fmt"
  "bufio"
	"os"
	"time"
  "strings"
)

type State int

const (
	INIT State = iota
	ACTIVE
	BREAK
  WAITING
	DONE
)

var activeTime int = 25
var breakTime int = 5
var intervals int = 3

func usage() {
	fmt.Println("Usage: pomo [command]")
}

func getStateRepr(s State) string {
	if s == 0 {
		return "Initial"
	} else if s == 1 {
		return "Active"
	} else if s == 2 {
		return "Break"
	} else if s == 3 {
    return "Waiting"
  }
	return "Done"
}

func getStateSymbol(s State) string {
	if s == 0 {
		return "🔵"
	} else if s == 1 {
		return "🟢"
	} else if s == 2 {
		return "🔴"
	} else if s == 3 {
    return "🟣"
  }
	return "⚫️"
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func timer(timeAmount int, curState State, interval int) {
	stateRepr := getStateRepr(curState)
  stateSymbol := getStateSymbol(curState)
	for t := 0; t < timeAmount; t++ {
		select {
		case <-time.After(1 * time.Second):
			clearScreen()
			fmt.Println("🍎 Time to focus")
			fmt.Printf("   State: %s %s\n", stateRepr, stateSymbol)
			fmt.Printf("   Interval: %d\n", interval)
			fmt.Printf("   Time Remaining: %ds\n", timeAmount-t)
		}
	}
}

func awaitUserRes(s State) {
  clearScreen()
  if s == ACTIVE {
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

func handleStartCommand(subcommands []string) {
	curState := INIT

	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	for i := 0; i < intervals; i++ {
		curState = ACTIVE
		timer(activeTime, curState, i)
    awaitUserRes(curState)
		curState = BREAK
		timer(breakTime, curState, i)
    awaitUserRes(curState)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	switch args[1] {
	case "start":
		handleStartCommand(args[2:])
	default:
		usage()
	}
}
